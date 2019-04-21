package postgre

import (
	"errors"
	"github.com/isa-mrs-tim6/Projekat/pkg/models"
	"time"
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

func (db *Store) GetRooms(id uint) ([]models.Room, error) {
	var retVal []models.Room
	if err := db.Set("gorm:auto_preload", true).Where("hotel_id = ?", id).Find(&retVal).Error; err != nil {
		return retVal, err
	}
	return retVal, nil
}

func (db *Store) AddRooms(id uint, rooms []models.Room) error {
	var retVal models.Hotel
	if err := db.Set("gorm:auto_preload", true).First(&retVal, id).Error; err != nil {
		return err
	}
	newRooms := append(retVal.Rooms, rooms...)
	retVal.Rooms = newRooms

	if err := db.Save(&retVal).Error; err != nil {
		return err
	}
	return nil
}

func (db *Store) DeleteRooms(id uint, rooms []models.Room) error {
	var retVal []models.HotelReservation
	if err := db.Set("gorm:auto_preload", true).Find(&retVal).Error; err != nil {
		return err
	}
	for i := 0; i < len(rooms); i++ {
		for _, v := range retVal {
			for _, room := range v.Rooms {
				if room.ID == rooms[i].ID && v.Occupation.End.After(time.Now()) {
					return errors.New("room is occupied and can't be deleted")
				}
			}
		}
	}

	for i := 0; i < len(rooms); i++ {
		db.Delete(&rooms[i])
	}
	return nil
}

func (db *Store) UpdateRoom(room models.Room) error {
	var retVal models.Room
	if err := db.Find(&retVal, room.ID).Error; err != nil {
		return err
	}
	retVal = room

	if err := db.Save(&retVal).Error; err != nil {
		return err
	}
	return nil
}
