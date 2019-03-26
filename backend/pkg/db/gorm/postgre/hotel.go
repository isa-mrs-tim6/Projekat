package postgre

import (
	"github.com/isa-mrs-tim6/Projekat/pkg/models"
)

type HotelDBInterface interface {
	GetHotels() ([]models.Hotel, error)
	GetHotelProfile(id uint) (models.Hotel, error)
	CreateHotel(hotel *models.Hotel) error
	UpdateHotel(hotel *models.Hotel) error
}

func (db *Store) GetHotels() ([]models.Hotel, error) {
	var retVal []models.Hotel
	if err := db.Set("gorm:auto_preload", true).Find(&retVal).Error; err != nil {
		return retVal, err
	}
	return retVal, nil
}

func (db *Store) GetHotelProfile(id uint) (models.HotelProfile, error) {
	var retVal models.Hotel
	if err := db.First(&retVal, id).Error; err != nil {
		return retVal.HotelProfile, err
	}
	return retVal.HotelProfile, nil
}

func (db *Store) CreateHotel(hotel *models.Hotel) error {
	if err := db.Create(&hotel).Error; err != nil {
		return err
	}
	return nil
}

func (db *Store) UpdateHotel(id uint, newProfile models.HotelProfile) error {
	var retVal models.Hotel
	if err := db.First(&retVal, id).Error; err != nil {
		return err
	}
	retVal.HotelProfile = newProfile

	if err := db.Save(&retVal).Error; err != nil {
		return err
	}
	return nil
}
