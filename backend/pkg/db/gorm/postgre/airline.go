package postgre

import (
	"github.com/isa-mrs-tim6/Projekat/pkg/models"
)

func (db *Store) GetAirlines() ([]models.Airline, error) {
	var retVal []models.Airline
	if err := db.Find(&retVal).Error; err != nil {
		return retVal, err
	}
	return retVal, nil
}

func (db *Store) CreateAirline(airline *models.Airline) error {
	if err := db.Create(&airline).Error; err != nil {
		return err
	}
	return nil
}

func (db *Store) GetAirlineProfile(id uint) (models.AirlineProfile, error) {
	var retVal models.Airline
	if err := db.First(&retVal, id).Error; err != nil {
		return retVal.AirlineProfile, err
	}
	return retVal.AirlineProfile, nil
}

func (db *Store) UpdateAirline(id uint, newProfile models.AirlineProfile) error {
	var retVal models.Airline
	if err := db.First(&retVal, id).Error; err != nil {
		return err
	}
	retVal.AirlineProfile = newProfile

	if err := db.Save(&retVal).Error; err != nil {
		return err
	}
	return nil
}
