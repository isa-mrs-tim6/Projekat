package postgre

import (
	"github.com/isa-mrs-tim6/Projekat/pkg/models"
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

func (db *Store) GetFlightQuickReservations(id uint) ([]models.FlightQuickReserveSeats, error) {
	var retVal []models.FlightQuickReserveSeats
	if err := db.Where("FlightID = FlightID").Find(&retVal).Error; err != nil {
		return retVal, err
	}
	return retVal, nil
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
	return retVal, nil
}
