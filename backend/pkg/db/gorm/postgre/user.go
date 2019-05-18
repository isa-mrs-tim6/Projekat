package postgre

import "github.com/isa-mrs-tim6/Projekat/pkg/models"

func (db *Store) GetUsers() ([]models.User, error) {
	var retVal []models.User
	if err := db.Set("gorm:auto_preload", true).Find(&retVal).Error; err != nil {
		return retVal, err
	}
	return retVal, nil
}

func (db *Store) UpdateUser(oldEmail string, params models.ProfileParams) error {
	var retVal models.User

	if err := db.Where("email = ?", oldEmail).First(&retVal).Error; err != nil {
		return err
	}

	retVal.Profile = params.Profile

	if err := db.Save(&retVal).Error; err != nil {
		return err
	}
	return nil
}

func (db *Store) GetUserProfile(email string) (models.Profile, error) {
	var retVal models.User
	if err := db.Where("email = ?", email).First(&retVal).Error; err != nil {
		return retVal.Profile, err
	}
	return retVal.Profile, nil
}

func (db *Store) GetUser(email string) (models.User, error) {
	var retVal models.User
	if err := db.Where("email = ?", email).First(&retVal).Error; err != nil {
		return retVal, err
	}
	return retVal, nil
}

func (db *Store) GetAirlineAdmin(email string) (models.AirlineAdmin, error) {
	var retVal models.AirlineAdmin
	if err := db.Where("email = ?", email).First(&retVal).Error; err != nil {
		return retVal, err
	}
	return retVal, nil
}

func (db *Store) GetRACAdmin(email string) (models.RentACarAdmin, error) {
	var retVal models.RentACarAdmin
	if err := db.Where("email = ?", email).First(&retVal).Error; err != nil {
		return retVal, err
	}
	return retVal, nil
}

func (db *Store) UpdateRACAdmin(id uint, newProfile models.Profile) error {
	var retVal models.RentACarAdmin
	if err := db.First(&retVal, id).Error; err != nil {
		return err
	}
	hashedPassword, err := createHash(retVal.Credentials.Password)
	if err != nil {
		return err
	}
	newProfile.Password = string(hashedPassword)

	retVal.Profile = newProfile
	retVal.FirstPassChanged = true
	if err := db.Save(&retVal).Error; err != nil {
		return err
	}
	return nil
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
	hashedPassword, err := createHash(retVal.Credentials.Password)
	if err != nil {
		return err
	}
	newProfile.Password = string(hashedPassword)

	retVal.Profile = newProfile
	retVal.FirstPassChanged = true
	if err := db.Save(&retVal).Error; err != nil {
		return err
	}
	return nil
}

func (db *Store) GetSystemAdmin(email string) (models.SystemAdmin, error) {
	var retVal models.SystemAdmin
	if err := db.Where("email = ?", email).First(&retVal).Error; err != nil {
		return retVal, err
	}
	return retVal, nil
}

func (db *Store) UpdateSystemAdmin(id uint, newProfile models.Profile) error {
	var retVal models.SystemAdmin
	if err := db.First(&retVal, id).Error; err != nil {
		return err
	}
	hashedPassword, err := createHash(retVal.Credentials.Password)
	if err != nil {
		return err
	}
	newProfile.Password = string(hashedPassword)

	retVal.Profile = newProfile
	retVal.FirstPassChanged = true
	if err := db.Save(&retVal).Error; err != nil {
		return err
	}
	return nil
}
