package postgre

import "github.com/isa-mrs-tim6/Projekat/pkg/models"

func (db *Store) GetUsers() ([]models.User, error) {
	var retVal []models.User
	if err := db.Set("gorm:auto_preload", true).Find(&retVal).Error; err != nil {
		return retVal, err
	}
	return retVal, nil
}

func (db *Store) UpdateUser(id uint, newProfile models.Profile) error {
	var retVal models.User
	if err := db.First(&retVal, id).Error; err != nil {
		return err
	}
	retVal.Profile = newProfile

	if err := db.Save(&retVal).Error; err != nil {
		return err
	}
	return nil
}

func (db *Store) GetUserProfile(id uint) (models.Profile, error) {
	var retVal models.User
	if err := db.First(&retVal, id).Error; err != nil {
		return retVal.Profile, err
	}
	return retVal.Profile, nil
}

func (db *Store) GetAirlineAdmin(email string) (models.AirlineAdmin, error) {
	var retVal models.AirlineAdmin
	if err := db.Where("email = ?", email).First(&retVal).Error; err != nil {
		return retVal, err
	}
	return retVal, nil
}

func (db *Store) GetHotelAdmin(email string) (models.HotelAdmin, error) {
	var retVal models.HotelAdmin
	if err := db.Where("email = ?", email).First(&retVal).Error; err != nil {
		return retVal, err
	}
	return retVal, nil
}

func (db *Store) UpdateHotelAdmin(id uint, newProfile models.Profile) error {
	var retVal models.HotelAdmin
	if err := db.First(&retVal, id).Error; err != nil {
		return err
	}
	retVal.Profile = newProfile

	if err := db.Save(&retVal).Error; err != nil {
		return err
	}
	return nil
}
