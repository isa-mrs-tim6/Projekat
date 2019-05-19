package postgre

import (
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
