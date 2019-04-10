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

func (db *Store) LoginAdmin(credentials models.Credentials) (string, error) {

	var aAdmin models.AirlineAdmin
	var hAdmin models.HotelAdmin
	var rAdmin models.RentACarAdmin
	var sAdmin models.SystemAdmin

	if err := db.First(&aAdmin, "email = ?", credentials.Email).Error; err == nil {
		if err := bcrypt.CompareHashAndPassword([]byte(aAdmin.Password), []byte(credentials.Password)); err != nil {
			return "", err
		}

		if !aAdmin.RegistrationComplete {
			return "", errors.New("not activated")
		}

		return "AirlineAdmin", nil
	}

	if err := db.First(&hAdmin, "email = ?", credentials.Email).Error; err == nil {
		if err := bcrypt.CompareHashAndPassword([]byte(hAdmin.Password), []byte(credentials.Password)); err != nil {
			return "", err
		}

		if !hAdmin.RegistrationComplete {
			return "", errors.New("not activated")
		}

		return "HotelAdmin", nil
	}

	if err := db.First(&rAdmin, "email = ?", credentials.Email).Error; err == nil {
		if err := bcrypt.CompareHashAndPassword([]byte(rAdmin.Password), []byte(credentials.Password)); err != nil {
			return "", err
		}

		if !rAdmin.RegistrationComplete {
			return "", errors.New("not activated")
		}

		return "Rent-A-CarAdmin", nil
	}

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
}
