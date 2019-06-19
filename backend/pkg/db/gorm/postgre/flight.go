package postgre

import (
	"github.com/isa-mrs-tim6/Projekat/pkg/models"
	"sort"
)

func (db *Store) GetAirplanes() ([]models.Airplane, error) {
	var retVal []models.Airplane
	if err := db.Set("gorm:auto_preload", true).Find(&retVal).Error; err != nil {
		return retVal, err
	}
	return retVal, nil
}

func (db *Store) UpdateFlight(id uint, newFlight models.Flight) error {
	var retVal models.Flight
	if err := db.First(&retVal, id).Error; err != nil {
		return err
	}
	retVal = newFlight

	if err := db.Save(&retVal).Error; err != nil {
		return err
	}
	return nil
}

func (db *Store) GetFlight(id uint) (models.Flight, error) {
	var retVal models.Flight
	if err := db.Preload("Origin").Preload("Destination").Preload("Airplane.Seats").
		First(&retVal, id).Error; err != nil {
		return retVal, err
	}
	return retVal, nil
}

func (db *Store) GetSeat(id uint) (models.Seat, error) {
	var retVal models.Seat
	if err := db.First(&retVal, id).Error; err != nil {
		return retVal, err
	}
	return retVal, nil
}

func (db *Store) GetFlightQuickReservations(id uint) ([]models.QuickFlightReservationGDTO, error) {
	var retVal []models.QuickFlightReservationGDTO
	if err := db.Table("flight_reservations").Select("seats.Number, flight_reservations.price, seats.reservation_id, reservations.id").
		Joins("INNER JOIN seats ON seats.reservation_id = flight_reservations.id").
		Joins("LEFT JOIN reservations ON flight_reservations.id = reservations.reservation_flight_id").
		Where("flight_reservations.is_quick_reserve = true AND flight_reservations.deleted_at IS NULL AND flight_reservations.flight_id = ?", id).Scan(&retVal).Error; err != nil {
		return retVal, err
	}
	return retVal, nil
}

func (db *Store) GetCompanyQuickReservations(id uint) ([]models.QuickFlightReservationGDTOV2, error) {
	var retVal []models.QuickFlightReservationGDTOV2
	if err := db.Table("flight_reservations").Select("seats.class, flight_reservations.price, flight_reservations.id, dest.name as DestName,origin.name as OriginName, flights.departure, flight_reservations.flight_id, " +
		"	flights.price_economy, flights.price_business, flights.price_firstclass").
		Joins("INNER JOIN seats ON seats.reservation_id = flight_reservations.id").
		Joins("INNER JOIN flights ON flight_reservations.flight_id = flights.id").
		Joins("INNER JOIN destinations as dest ON flights.destination_id = dest.id").
		Joins("INNER JOIN destinations as origin ON flights.origin_id = origin.id").
		Joins("LEFT JOIN reservations ON flight_reservations.id = reservations.reservation_flight_id").
		Where("flight_reservations.is_quick_reserve = true AND flight_reservations.deleted_at IS NULL AND flights.airline_id = ? AND reservations.reservation_flight_id is NULL", id).Scan(&retVal).Error; err != nil {
		return retVal, err
	}
	return retVal, nil
}

func (db *Store) CreateFlightQuickReservation(reservation *models.FlightReservation) error {
	if err := db.Create(&reservation).Error; err != nil {
		return err
	}
	return nil
}

func (db *Store) RemoveFLightQuickReservation(resDTO models.QuickFlightReservationGDTO) error {
	if err := db.Table("flight_reservations").Where("id = ?", resDTO.ReservationID).
		Delete(models.FlightReservation{}).Error; err != nil {
		return err
	}
	if err := db.Table("seats").Where("reservation_id = ?", resDTO.ReservationID).
		Update("reservation_id",0).Error; err != nil{
			return err
	}
	return nil
}

func (db *Store) GetCompanyAirplanes(id uint) ([]models.Airplane, error) {
	var retVal []models.Airplane
	if err := db.Set("gorm:auto_preload", true).Where("airline_id = ?", id).Find(&retVal).Error; err != nil {
		return retVal, err
	}
	return retVal, nil
}

func (db *Store) GetAirplane(id uint, name string) (models.Airplane, error) {
	var retVal models.Airplane
	if err := db.Where("airline_id = ? AND name = ?", id, name).Find(&retVal).Error; err != nil {
		return retVal, err
	}
	return retVal, nil
}

func (db *Store) CreateFlight(flight *models.Flight) error {
	if err := db.Create(&flight).Error; err != nil {
		return err
	}
	return nil
}

func (db *Store) GetCompanyFlights(id uint) ([]models.Flight, error) {
	var retVal []models.Flight
	if err := db.Preload("Origin").Preload("Destination").Preload("Airplane").
		Preload("Airplane.Seats").Where("airline_id = ?", id).Find(&retVal).Error; err != nil {
		return retVal, err
	}
	for _, element := range retVal {
		sort.SliceStable(element.Airplane.Seats, func(i, j int) bool {
			return element.Airplane.Seats[i].Number < element.Airplane.Seats[j].Number
		})
	}
	return retVal, nil
}
