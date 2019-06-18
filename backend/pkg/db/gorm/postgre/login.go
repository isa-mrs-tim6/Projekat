package postgre

import (
	"errors"
	"github.com/isa-mrs-tim6/Projekat/pkg/models"
	"golang.org/x/crypto/bcrypt"
)

func (db *Store) LoginUser(credentials models.Credentials) error {

	var user models.User

	if err := db.First(&user, "email = ?", credentials.Email).Error; err != nil {
		return err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(credentials.Password)); err != nil {
		return err
	}

	if !user.RegistrationComplete {
		return errors.New("not activated")
	}

	return nil
}

func (db *Store) LoginAdmin(credentials models.Credentials, accountType string) (string, error) {

	switch accountType {
	case "SystemAdmin":
		var sAdmin models.SystemAdmin
		if err := db.First(&sAdmin, "email = ?", credentials.Email).Error; err == nil {
			if err := bcrypt.CompareHashAndPassword([]byte(sAdmin.Password), []byte(credentials.Password)); err != nil {
				return "", err
			}

			if !sAdmin.RegistrationComplete {
				return "", errors.New("not activated")
			}

			return "SystemAdmin", nil
		}

		return "", nil
	case "AirlineAdmin":
		var aAdmin models.AirlineAdmin
		if err := db.First(&aAdmin, "email = ?", credentials.Email).Error; err == nil {
			if err := bcrypt.CompareHashAndPassword([]byte(aAdmin.Password), []byte(credentials.Password)); err != nil {
				return "", err
			}

			if !aAdmin.RegistrationComplete {
				return "", errors.New("not activated")
			}

			return "AirlineAdmin", nil
		}
	case "HotelAdmin":
		var hAdmin models.HotelAdmin
		if err := db.First(&hAdmin, "email = ?", credentials.Email).Error; err == nil {
			if err := bcrypt.CompareHashAndPassword([]byte(hAdmin.Password), []byte(credentials.Password)); err != nil {
				return "", err
			}

			if !hAdmin.RegistrationComplete {
				return "", errors.New("not activated")
			}

			return "HotelAdmin", nil
		}
	case "Rent-A-CarAdmin":
		var rAdmin models.RentACarAdmin
		if err := db.First(&rAdmin, "email = ?", credentials.Email).Error; err == nil {
			if err := bcrypt.CompareHashAndPassword([]byte(rAdmin.Password), []byte(credentials.Password)); err != nil {
				return "", err
			}

			if !rAdmin.RegistrationComplete {
				return "", errors.New("not activated")
			}

			return "Rent-A-CarAdmin", nil
		}
	}
	return "", errors.New("unknown type")
}

func (db *Store) CheckFirstPassChanged(email string, accountType string) (bool, error) {

	switch accountType {
	case "SystemAdmin":
		var retVal models.SystemAdmin
		if err := db.Where("email = ?", email).First(&retVal).Error; err != nil {
			return false, err
		}
		return retVal.FirstPassChanged, nil
	case "AirlineAdmin":
		var retVal models.AirlineAdmin
		if err := db.Where("email = ?", email).First(&retVal).Error; err != nil {
			return false, err
		}
		return retVal.FirstPassChanged, nil
	case "HotelAdmin":
		var retVal models.HotelAdmin
		if err := db.Where("email = ?", email).First(&retVal).Error; err != nil {
			return false, err
		}
		return retVal.FirstPassChanged, nil
	case "Rent-A-CarAdmin":
		var retVal models.RentACarAdmin
		if err := db.Where("email = ?", email).First(&retVal).Error; err != nil {
			return false, err
		}
		return retVal.FirstPassChanged, nil
	}
	return false, nil
}
