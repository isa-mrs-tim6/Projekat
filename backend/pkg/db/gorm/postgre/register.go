package postgre

import (
	"github.com/isa-mrs-tim6/Projekat/pkg/models"
	"golang.org/x/crypto/bcrypt"
)

type RegisterInterface interface {
	RegisterSystemAdmin(credentials models.Credentials) error
	CompleteRegistrationSystemAdmin(hashedEmail string) error
}

func (db *Store) CompleteRegistration(email string, accountType string) error {
	switch accountType {
	case "SystemAdmin":
		var retVal models.SystemAdmin
		if err := db.Model(&retVal).Where("email = ?", email).Update("registration_complete", true).Error; err != nil {
			return err
		}
		return nil
	case "AirlineAdmin":
		var retVal models.AirlineAdmin
		if err := db.Model(&retVal).Where("email = ?", email).Update("registration_complete", true).Error; err != nil {
			return err
		}
		return nil
	case "HotelAdmin":
		var retVal models.HotelAdmin
		if err := db.Model(&retVal).Where("email = ?", email).Update("registration_complete", true).Error; err != nil {
			return err
		}
		return nil
	case "Rent-A-CarAdmin":
		var retVal models.RentACarAdmin
		if err := db.Model(&retVal).Where("email = ?", email).Update("registration_complete", true).Error; err != nil {
			return err
		}
		return nil
	case "User":
		var retVal models.User
		if err := db.Model(&retVal).Where("email = ?", email).Update("registration_complete", true).Error; err != nil {
			return err
		}
		return nil
	}
	return nil
}

func (db *Store) RegisterAdmin(credentials models.Credentials, accountType string, companyID uint) error {
	hashedPassword, err := createHash(credentials.Password)
	if err != nil {
		return err
	}

	switch accountType {
	case "SystemAdmin":
		var admin models.SystemAdmin
		admin.Credentials = credentials
		admin.Password = string(hashedPassword)
		if err := db.Create(&admin).Error; err != nil {
			return err
		}
		return nil
	case "AirlineAdmin":
		var admin models.AirlineAdmin
		admin.Credentials = credentials
		admin.Password = string(hashedPassword)
		admin.AirlineID = companyID
		if err := db.Create(&admin).Error; err != nil {
			return err
		}
		return nil
	case "HotelAdmin":
		var admin models.HotelAdmin
		admin.Credentials = credentials
		admin.Password = string(hashedPassword)
		admin.HotelID = companyID
		if err := db.Create(&admin).Error; err != nil {
			return err
		}
		return nil
	case "Rent-A-CarAdmin":
		var admin models.RentACarAdmin
		admin.Credentials = credentials
		admin.Password = string(hashedPassword)
		admin.RentACarCompanyID = companyID
		if err := db.Create(&admin).Error; err != nil {
			return err
		}
		return nil
	}
	return nil
}

func (db *Store) RegisterUser(user models.User) error {
	hashedPassword, err := createHash(user.Credentials.Password)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)
	//user.RegistrationComplete = false

	if err := db.Create(&user).Error; err != nil {
		return err
	}
	return nil
}

func createHash(hashObject string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(hashObject), bcrypt.DefaultCost)
}
