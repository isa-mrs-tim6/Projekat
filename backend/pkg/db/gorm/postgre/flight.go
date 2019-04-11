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

func (db *Store) CreateFlight(flight *models.Flight) error {
	if err := db.Create(&flight).Error; err != nil {
		return err
	}
	return nil
}
