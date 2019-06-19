package postgre

import (
	"errors"
	"fmt"
	"github.com/isa-mrs-tim6/Projekat/pkg/models"
	"strconv"
	"strings"
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

	var invitedBy models.User
	var reservation models.Reservation

	for _, res := range masters {
		if res.MasterRef != 0 {
			db.Where("id = ?", res.MasterRef).
				First(&reservation)

			db.Table("users").
				Where("id = ?", reservation.UserID).
				First(&invitedBy)
		}

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
		reservations = append(reservations, models.ReservationDAO{Master: res, Slaves: slaves, InvitedBy: invitedBy})
	}

	return reservations, nil
}

func (db *Store) ReserveVehicle(masterRef uint, params models.VehicleReservationParams, userID uint) error {
	var reservation models.RentACarReservation
	var vehicle models.Vehicle
	var location models.Location

	tx := db.Begin()

	reservation.CompanyID = params.CompanyID

	if err := tx.Raw("SELECT * FROM vehicles WHERE id = ? FOR UPDATE", params.VehicleID).Scan(&vehicle).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Where("id=?", params.LocationID).First(&location).Error; err != nil {
		tx.Rollback()
		return err
	}

	reservation.Vehicle = vehicle
	reservation.Price = db.CalculatePriceVehicle(userID, params.Price)
	reservation.Location = location.Address.Address

	start, _ := strconv.ParseInt(params.StartDate, 10, 64)
	end, _ := strconv.ParseInt(params.EndDate, 10, 64)

	startDate := time.Unix(0, start*int64(time.Millisecond))
	endDate := time.Unix(0, end*int64(time.Millisecond))

	reservation.Occupation.Beginning = startDate
	reservation.Occupation.End = endDate

	// Grab master reservation
	var masterReservation models.Reservation

	if err := tx.First(&masterReservation, masterRef).Error; err != nil {
		tx.Rollback()
		return err
	}

	var racReservations []models.RentACarReservation
	if err := tx.Where("vehicle_id = ?", vehicle.ID).Find(&racReservations).Error; err != nil {
		tx.Rollback()
		return err
	}
	for _, res := range racReservations {
		if !(res.Occupation.Beginning.After(endDate) ||
			res.Occupation.End.Before(startDate)) {
			tx.Rollback()
			return errors.New("vehicle taken")
		}
	}

	masterReservation.ReservationRentACar = reservation
	masterReservation.ReservationRentACarID = reservation.ID

	if err := tx.Save(&masterReservation).Error; err != nil {
		tx.Rollback()
		return err
	}

	// Get all associated reservations
	var reservations []models.Reservation
	if err := tx.Where("master_ref = ?", masterRef).Find(&reservations).Error; err != nil {
		tx.Rollback()
		return err
	}

	for _, associatedReservation := range reservations {
		reservation := models.RentACarReservation{
			CompanyID: params.CompanyID,
			Vehicle:   vehicle,
			Price:     db.CalculatePriceVehicle(userID, params.Price),
			Location:  location.Address.Address,
			Occupation: models.Occupation{
				Beginning: startDate,
				End:       endDate,
			},
		}
		associatedReservation.ReservationRentACar = reservation
		associatedReservation.ReservationRentACarID = reservation.ID
		if err := tx.Save(&associatedReservation).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	tx.Commit()

	return nil
}

func (db *Store) GetPriceScale(userID uint) (float64, uint, error) {
	var numOfReservations uint
	var reward models.ReservationReward
	db.Table("reservations").Where("user_id = ? AND deleted_at IS NULL", userID).Count(&numOfReservations)
	db.First(&reward, "required_number < ?", numOfReservations)
	return reward.PriceScale, numOfReservations, nil
}

func (db *Store) CalculatePriceVehicle(userID uint, originalPrice float64) float64 {
	price := float64(0.0)

	var numOfReservations uint
	var reward models.ReservationReward
	db.Table("reservations").Where("user_id = ?", userID).Count(&numOfReservations)
	db.First(&reward, "required_number < ?", numOfReservations)
	if reward.PriceScale != 0 {
		price *= reward.PriceScale
	}
	return price
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

	var expireTime time.Time
	var isExpiring bool

	for i := 1; i < len(params.Users); i++ {
		var friend models.User
		if params.Users[i].Email != "" { // Registered user
			db.Where("email = ?", params.Users[i].Email).First(&friend)
			params.Users[i].ID = friend.ID
			params.Users[i].Name = friend.Name
			params.Users[i].Surname = friend.Surname
			params.Users[i].Passport = friend.Passport
			expireTime = time.Now().AddDate(0, 0, 3)
			isExpiring = true
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
			MasterRef:  masterReservation.ID,
			IsExpiring: isExpiring,
			ExpireTime: expireTime,
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
	params models.HotelReservationParams) (uint, string, string, string, string, string, string, string, error) {
	// Check parameters
	if params.Rooms == nil || len(params.Rooms) == 0 {
		return 0, "", "", "", "", "", "", "", errors.New("invalid parameters")
	}

	var userStr string
	var featurestrSplice []string
	var roomsStrSplice []string
	var hotelStr string

	for _, v := range params.Rooms {
		roomsStrSplice = append(roomsStrSplice, strconv.Itoa(v.Number))
	}

	for _, v := range params.Features {
		featurestrSplice = append(featurestrSplice, v.Name)
	}

	tx := db.Begin()

	// Grab master reservation
	var masterReservation models.Reservation
	if err := tx.First(&masterReservation, masterReservationID).Error; err != nil {
		tx.Rollback()
		return 0, "", "", "", "", "", "", "", err
	}

	var hReservations []models.HotelReservation
	var ids []uint

	for _, room := range params.Rooms {
		ids = append(ids, room.ID)
	}

	var rooms []models.Room
	if err := tx.Raw("SELECT * FROM rooms WHERE id in (?) FOR UPDATE", ids).Scan(&rooms).Error; err != nil {
		tx.Rollback()
		return 0, "", "", "", "", "", "", "", err
	}

	if err := tx.Table("hotel_reservations").
		Joins("JOIN room_reservations on room_reservations.hotel_reservation_id = hotel_reservations.id").
		Where("room_reservations.room_id in (?) AND"+
			"((hotel_reservations.beginning BETWEEN ? AND ?) OR (hotel_reservations.end BETWEEN ? AND ?))",
			ids, params.From, params.To, params.From, params.To).
		Find(&hReservations).
		Error; err != nil {
		tx.Rollback()
		return 0, "", "", "", "", "", "", "", err
	}

	if hReservations != nil && len(hReservations) != 0 {
		tx.Rollback()
		return 0, "", "", "", "", "", "", "", errors.New("Room taken")
	}

	var hotel models.Hotel
	tx.Find(&hotel, hotelID)
	hotelStr = hotel.Name + ", " + hotel.Address.Address

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
		Features: params.Features,
		Price:    db.CalculateHotelReservationPrice(userID, params.Rooms, params.Features, params.IsQuickReserve),
	}
	masterReservation.ReservationHotelID = hotelReservation.ID
	masterReservation.ReservationHotel = hotelReservation
	userStr = masterReservation.Name + " " + masterReservation.Surname

	//db.Create(&hotelReservation)
	if err := tx.Save(&masterReservation).Error; err != nil {
		tx.Rollback()
		return 0, "", "", "", "", "", "", "", err
	}

	// Get all associated reservations
	var reservations []models.Reservation
	if err := tx.Where("master_ref = ?", masterReservationID).Find(&reservations).Error; err != nil {
		tx.Rollback()
		return 0, "", "", "", "", "", "", "", err
	}

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
			Features: params.Features,
			Price:    db.CalculateHotelReservationPrice(userID, params.Rooms, params.Features, params.IsQuickReserve),
		}
		reservation.ReservationHotel = hotelReservation
		reservation.ReservationHotelID = hotelReservation.ID
		//db.Create(&hotelReservation)
		if err := tx.Save(&reservation).Error; err != nil {
			tx.Rollback()
			return 0, "", "", "", "", "", "", "", err
		}
	}

	tx.Commit()

	roomsStr := strings.Join(roomsStrSplice, ", ")
	featuresStr := strings.Join(featurestrSplice, ", ")
	price := fmt.Sprintf("%f", hotelReservation.Price)
	from := params.From.String()
	to := params.To.String()

	return masterReservation.ID, hotelStr, userStr, roomsStr, featuresStr, price, from, to, nil
}

func (db *Store) CalculateHotelReservationPrice(userID uint, sent []models.Room, features []*models.Feature,
	isQuickReserve bool) float64 {
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

	for _, v := range features {
		price += v.Price
	}

	if isQuickReserve {
		price *= 0.9
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

func (db *Store) CancelFlight(resID uint) error {
	var master models.Reservation
	var slaves []models.Reservation

	db.Where("id = ?", resID).
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
		First(&master)
	db.Table("seats").Where("reservation_id = ?", master.ReservationFlight.Seat.ReservationID).
		Update("reservation_id", 0)

	if !master.ReservationFlight.IsQuickReserve {
		db.Where("id=?", master.ID).Delete(master.ReservationFlight)
	}
	if !master.ReservationHotel.IsQuickReserve {
		db.Where("id=?", master.ID).Delete(master.ReservationHotel)
	}
	if !master.ReservationRentACar.IsQuickReserve {
		db.Where("id=?", master.ID).Delete(master.ReservationRentACar)
	}

	if master.MasterRef != 0 {
		db.Delete(&master)
		return nil
	}

	db.Where("master_ref = ?", master.ID).
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

	for _, slave := range slaves {
		db.Table("seats").Where("reservation_id = ?", slave.ReservationFlight.Seat.ReservationID).
			Update("reservation_id", 0)

		if !slave.ReservationFlight.IsQuickReserve {
			db.Where("id=?", slave.ID).Delete(slave.ReservationFlight)
		}
		if !slave.ReservationHotel.IsQuickReserve {
			db.Where("id=?", slave.ID).Delete(slave.ReservationHotel)
		}
		if !slave.ReservationRentACar.IsQuickReserve {
			db.Where("id=?", slave.ID).Delete(slave.ReservationRentACar)
		}
		db.Delete(slave)
	}

	db.Delete(&master)
	return nil
}

func (db *Store) CancelHotel(resID uint) error {
	var master models.Reservation
	var slaves []models.Reservation

	db.Where("id = ?", resID).
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
		First(&master)

	if master.MasterRef != 0 {
		return nil
	}

	if !master.ReservationHotel.IsQuickReserve {
		db.Table("hotel_reservations").Where("id = ?", master.ReservationHotelID).
			Delete(master.ReservationHotel)
	}

	db.Table("reservations").Where("id=?", master.ID).
		Update("reservation_hotel_id", 0)

	db.Where("master_ref = ?", master.ID).
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

	for _, slave := range slaves {
		if !slave.ReservationHotel.IsQuickReserve {
			db.Table("hotel_reservations").Where("id = ?", slave.ReservationHotelID).
				Delete(slave.ReservationHotel)
		}

		db.Table("reservations").Where("id=?", slave.ID).
			Update("reservation_hotel_id", 0)
	}
	return nil
}

func (db *Store) CancelVehicle(resID uint) error {
	var master models.Reservation
	var slaves []models.Reservation

	db.Where("id = ?", resID).
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
		First(&master)

	if master.MasterRef != 0 {
		return nil
	}

	if !master.ReservationRentACar.IsQuickReserve {
		db.Table("rent_a_car_reservations").Where("id=?", master.ReservationRentACarID).
			Delete(master.ReservationRentACar)
	}

	db.Table("reservations").Where("id=?", master.ID).
		Update("reservation_rent_a_car_id", 0)

	db.Where("master_ref = ?", master.ID).
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

	for _, slave := range slaves {
		if !slave.ReservationRentACar.IsQuickReserve {
			db.Table("rent_a_car_reservations").Where("id=?", slave.ReservationRentACarID).
				Delete(slave.ReservationRentACar)
		}

		db.Table("reservations").Where("id=?", slave.ID).
			Update("reservation_rent_a_car_id", 0)
	}
	return nil
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
	var retVal []models.RentACarReservation
	var masters []models.Reservation

	db.Preload("ReservationRentACar").Find(&masters)

	start, _ := strconv.ParseInt(params.StartDate, 10, 64)
	end, _ := strconv.ParseInt(params.EndDate, 10, 64)

	startDate := time.Unix(0, start*int64(time.Millisecond))
	endDate := time.Unix(0, end*int64(time.Millisecond))

	if err := db.Preload("Vehicle").
		Where("company_id = ? AND is_quick_reserve = true AND beginning = ? AND rent_a_car_reservations.end = ?",
			params.CompanyID, startDate, endDate).
		Find(&reservations).Error; err != nil {
		return retVal, err
	}

	var found bool

	for _, res := range reservations {
		found = false
		for _, master := range masters {
			if master.ReservationRentACarID == res.ID {
				found = true
				break
			}
		}
		if !found {
			retVal = append(retVal, res)
		}
	}

	return retVal, nil
}

func (db *Store) CompleteQuickResVehicle(params models.CompleteQuickResVehParams) error {
	var master models.Reservation
	var slave models.RentACarReservation
	var loc models.Location

	if err := db.Preload("ReservationRentACar").
		Where("id = ?", params.MasterID).First(&master).Error; err != nil {
		return err
	}
	if err := db.Where("id = ?", params.ReservationID).First(&slave).Error; err != nil {
		return err
	}

	if err := db.Where("id = ?", params.LocationID).First(&loc).Error; err != nil {
		return err
	}

	master.ReservationRentACar = slave
	master.ReservationRentACarID = slave.ID
	master.ReservationRentACar.Location = loc.Address.Address

	if err := db.Save(&master).Error; err != nil {
		return err
	}
	return nil
}
