package postgre

import (
	"errors"
	"github.com/isa-mrs-tim6/Projekat/pkg/models"
	"strconv"
	"time"
)

func (db *Store) GetReservationRewards() ([]models.ReservationReward, error) {
	var retVal []models.ReservationReward
	if err := db.Set("gorm:auto_preload", true).Find(&retVal).Error; err != nil {
		return retVal, err
	}
	return retVal, nil
}

func (db *Store) UpdateReservationRewards(rewards []models.ReservationReward) {
	db.Delete(models.ReservationReward{})
	for _, v := range rewards {
		db.Create(v)
	}
}

func (db *Store) GetReservation(reservationID uint) (models.Reservation, error) {
	var retVal models.Reservation
	if err := db.First(&retVal, "id = ?", reservationID).Error; err != nil {
		return retVal, err
	}
	return retVal, nil
}

func (db *Store) GetReservationGraphData(id uint) ([]models.ReservationGraphData, error) {
	var retVal []models.ReservationGraphData
	if err := db.Table("flight_reservations").Select("flight_reservations.id, flight_reservations.price, flights.departure").
		Joins("JOIN flights ON flights.id = flight_reservations.flight_id").Where("flights.airline_id = ?", id).Scan(&retVal).
		Error; err != nil {
		return nil, err
	}
	return retVal, nil
}
func (db *Store) GetUserReservations(id uint) ([]models.ReservationDAO, error) {
	var reservations []models.ReservationDAO
	var masters []models.Reservation
	var slaves []models.Reservation

	if err := db.Order("created_at desc").
		Where("user_id = ?", id).
		Preload("ReservationFlight.Flight.Origin").
		Preload("ReservationFlight.Flight.Destination").
		Preload("ReservationFlight.Flight.Layovers").
		Preload("ReservationFlight.Seat").
		Preload("ReservationFlight.Features").
		Preload("ReservationHotel.Rooms").
		Preload("ReservationHotel.Ratings").
		Preload("ReservationHotel.Features").
		Preload("ReservationHotel.Hotel").
		Preload("ReservationRentACar.RentACarCompany").
		Preload("ReservationRentACar.Vehicle").
		Find(&masters).Error; err != nil {
		return nil, err
	}

	for _, res := range masters {
		db.Where("master_ref = ?", res.ID).
			Preload("ReservationFlight.Flight.Origin").
			Preload("ReservationFlight.Flight.Destination").
			Preload("ReservationFlight.Flight.Layovers").
			Preload("ReservationFlight.Seat").
			Preload("ReservationFlight.Features").
			Preload("ReservationHotel.Rooms").
			Preload("ReservationHotel.Ratings").
			Preload("ReservationHotel.Features").
			Preload("ReservationHotel.Hotel").
			Preload("ReservationRentACar.RentACarCompany").
			Preload("ReservationRentACar.Vehicle").
			Find(&slaves)
		reservations = append(reservations, models.ReservationDAO{Master: res, Slaves: slaves})
	}

	return reservations, nil
}

func (db *Store) ReserveVehicle(masterRef uint, params models.VehicleReservationParams) error {
	var reservation models.RentACarReservation
	var vehicle models.Vehicle
	var location models.Location

	reservation.CompanyID = params.CompanyID

	if err := db.Where("id=?", params.VehicleID).First(&vehicle).Error; err != nil {
		return err
	}

	if err := db.Where("id=?", params.LocationID).First(&location).Error; err != nil {
		return err
	}

	reservation.Vehicle = vehicle
	reservation.Price = params.Price
	reservation.Location = location.Address.Address

	start, _ := strconv.ParseInt(params.StartDate, 10, 64)
	end, _ := strconv.ParseInt(params.EndDate, 10, 64)

	startDate := time.Unix(0, start*int64(time.Millisecond))
	endDate := time.Unix(0, end*int64(time.Millisecond))

	reservation.Occupation.Beginning = startDate
	reservation.Occupation.End = endDate

	// Grab master reservation
	var masterReservation models.Reservation
	db.First(&masterReservation, masterRef)

	masterReservation.ReservationRentACar = reservation
	masterReservation.ReservationRentACarID = reservation.ID
	db.Save(&masterReservation)

	// Get all associated reservations
	var reservations []models.Reservation
	db.Where("master_ref = ?", masterRef).Find(&reservations)

	for _, associated_reservation := range reservations {
		reservation := models.RentACarReservation{
			CompanyID: params.CompanyID,
			Vehicle:   vehicle,
			Price:     params.Price,
			Location:  location.Address.Address,
			Occupation: models.Occupation{
				Beginning: startDate,
				End:       endDate,
			},
		}
		associated_reservation.ReservationRentACar = reservation
		associated_reservation.ReservationRentACarID = reservation.ID
		db.Save(&associated_reservation)
	}

	return nil
}

func (db *Store) ReserveFlight(flightID uint64, params models.FlightReservationParams) (uint, error) {
	// Check parameters
	if params.Users == nil || params.Seats == nil || len(params.Users) != len(params.Seats) || len(params.Users) == 0 {
		return 0, errors.New("invalid parameters")
	}

	// Create master reservation, { Users[0], Seats[0] } combination
	masterReservation := models.Reservation{
		Passenger: models.Passenger{
			UserID:   params.Users[0].ID,
			UserInfo: params.Users[0].UserInfo,
		},
		ReservationFlight: models.FlightReservation{
			Seat:           &params.Seats[0],
			FlightID:       uint(flightID),
			FlightRating:   0,
			CompanyRating:  0,
			IsQuickReserve: params.IsQuickReserve,
			Price:          db.CalculatePriceFlight(uint(flightID), params.Seats[0].ID, params.Users[0].ID, nil, params.IsQuickReserve), // TODO Add support for features
			Features:       nil,                                                                                                         // TODO Add support for features
		},
	}
	db.Create(&masterReservation)

	for i := 1; i < len(params.Users); i++ {
		var friend models.User
		if params.Users[i].Email != "" { // Registered user
			db.Where("email = ?", params.Users[i].Email).First(&friend)
			params.Users[i].ID = friend.ID
			params.Users[i].Name = friend.Name
			params.Users[i].Surname = friend.Surname
			params.Users[i].Passport = friend.Passport
		}

		reservation := models.Reservation{
			Passenger: models.Passenger{
				UserID:   params.Users[i].ID,
				UserInfo: params.Users[i].UserInfo,
			},
			ReservationFlight: models.FlightReservation{
				Seat:           &params.Seats[i],
				FlightID:       uint(flightID),
				FlightRating:   0,
				CompanyRating:  0,
				IsQuickReserve: params.IsQuickReserve,
				Price:          db.CalculatePriceFlight(uint(flightID), params.Seats[i].ID, params.Users[i].ID, nil, params.IsQuickReserve), // TODO Add support for features
				Features:       nil,                                                                                                         // TODO Add support for features
			},
			MasterRef: masterReservation.ID,
		}
		db.Create(&reservation)
	}

	return masterReservation.ID, nil
}

func (db *Store) CalculatePriceFlight(flightID uint, seatID uint, userID uint, features []models.FeatureAirline, isQuickReserve bool) float64 {
	var flightPrices models.PriceList
	var seat models.Seat
	price := float64(0.0)

	// Add seat price
	db.Table("flights").Select("price_economy, price_business, price_firstclass,"+
		"small_suitcase, big_suitcase, quick_reservation_price_scale").Where("id = ?", flightID).Scan(&flightPrices)
	db.Table("seats").Select("class").Where("id = ?", seatID).Scan(&seat)
	switch seat.Class {
	case "ECONOMIC":
		price += flightPrices.PriceECONOMY
	case "BUSINESS":
		price += flightPrices.PriceBUSINESS
	case "FIRST":
		price += flightPrices.PriceFIRSTCLASS
	}
	if isQuickReserve {
		price *= flightPrices.QuickReservationPriceScale
	}

	// Add prices of features
	if features != nil {
		for _, feature := range features {
			price += feature.Price
		}
	}

	// Add discount based on number of reservations
	var numOfReservations uint
	var reward models.ReservationReward
	db.Table("reservations").Where("user_id = ?", userID).Count(&numOfReservations)
	db.First(&reward, "required_number < ?", numOfReservations)
	if reward.PriceScale != 0 {
		price *= reward.PriceScale
	}

	return price
}

func (db *Store) ReserveHotel(masterReservationID uint, hotelID uint, userID uint,
	params models.HotelReservationParams) (uint, error) {
	// Check parameters
	if params.Rooms == nil || len(params.Rooms) == 0 {
		return 0, errors.New("invalid parameters")
	}

	// Grab master reservation
	var masterReservation models.Reservation
	db.First(&masterReservation, masterReservationID)

	// Create hotel reservation
	hotelReservation := models.HotelReservation{
		Rooms:          params.Rooms,
		HotelID:        hotelID,
		IsQuickReserve: params.IsQuickReserve,
		Occupation: models.Occupation{
			Beginning: params.From,
			End:       params.To,
		},
		Ratings:  nil,
		Features: nil,
		Price:    db.CalculateHotelReservationPrice(userID, params.Rooms, params.IsQuickReserve),
	}
	masterReservation.ReservationHotelID = hotelReservation.ID
	masterReservation.ReservationHotel = hotelReservation
	//db.Create(&hotelReservation)
	db.Save(&masterReservation)

	// Get all associated reservations
	var reservations []models.Reservation
	db.Where("master_ref = ?", masterReservationID).Find(&reservations)

	for _, reservation := range reservations {
		hotelReservation := models.HotelReservation{
			Rooms:          params.Rooms,
			HotelID:        hotelID,
			IsQuickReserve: params.IsQuickReserve,
			Occupation: models.Occupation{
				Beginning: params.From,
				End:       params.To,
			},
			Ratings:  nil,
			Features: nil,
			Price:    db.CalculateHotelReservationPrice(userID, params.Rooms, params.IsQuickReserve),
		}
		reservation.ReservationHotel = hotelReservation
		reservation.ReservationHotelID = hotelReservation.ID
		//db.Create(&hotelReservation)
		db.Save(&reservation)
	}

	return masterReservation.ID, nil
}

func (db *Store) CalculateHotelReservationPrice(userID uint, sent []models.Room, isQuickReserve bool) float64 {
	var foundRooms []models.Room
	var foundRoomIds []uint
	price := float64(0.0)

	for _, v := range sent {
		foundRoomIds = append(foundRoomIds, v.ID)
	}
	db.Where(foundRoomIds).Find(&foundRooms)

	// Add room prices
	for _, v := range foundRooms {
		price += v.Price
	}

	// Add discount based on number of reservations
	var numOfReservations uint
	var reward models.ReservationReward
	db.Table("reservations").Where("user_id = ?", userID).Count(&numOfReservations)
	db.First(&reward, "required_number < ?", numOfReservations)
	if reward.PriceScale != 0 {
		price *= reward.PriceScale
	}

	return price
}

func (db *Store) GetQuickVehRes(id uint) ([]models.RentACarReservation, error) {
	var reservations []models.RentACarReservation

	if err := db.Where("vehicle_id = ? AND is_quick_reserve = true", id).Find(&reservations).Error; err != nil {
		return reservations, err
	}

	return reservations, nil
}

func (db *Store) GetCompanyQuickVehicle(params models.VehicleQuickResParams) ([]models.RentACarReservation, error) {
	var reservations []models.RentACarReservation
	start, _ := strconv.ParseInt(params.StartDate, 10, 64)
	end, _ := strconv.ParseInt(params.EndDate, 10, 64)

	startDate := time.Unix(0, start*int64(time.Millisecond))
	endDate := time.Unix(0, end*int64(time.Millisecond))

	if err := db.Preload("Vehicle").
		Where("company_id = ? AND is_quick_reserve = true AND beginning >= ? AND rent_a_car_reservations.end <= ?",
			params.CompanyID, startDate, endDate).
		Find(&reservations).Error; err != nil {
		return reservations, err
	}

	return reservations, nil
}

func (db *Store) CompleteQuickResVehicle(params models.CompleteQuickResVehParams) error {
	return nil
}
