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

func (db *Store) GetReservationGraphData(id uint) ([]models.ReservationGraphData, error) {
	var retVal []models.ReservationGraphData
	if err := db.Table("flight_reservations").Select("flight_reservations.id, flight_reservations.price, flights.departure").
		Joins("JOIN flights ON flights.id = flight_reservations.flight_id").Where("flights.airline_id = ?", id).Scan(&retVal).
		Error; err != nil {
		return nil, err
	}
	return retVal, nil
}

func (db *Store) ReserveVehicle(params models.VehicleReservationParams) error {
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

	if err := db.Create(&reservation).Error; err != nil {
		return err
	}

	return nil
}

func (db *Store) ReserveFlight(flightID uint64, params models.FlightReservationParams) error {
	// Check parameters
	if params.Users == nil || params.Seats == nil || len(params.Users) != len(params.Seats) || len(params.Users) == 0 {
		return errors.New("invalid parameters")
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

	return nil
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
