package postgre

import (
	"github.com/isa-mrs-tim6/Projekat/pkg/models"
	"golang.org/x/crypto/bcrypt"
)

type RegisterInterface interface {
	RegisterSystemAdmin(credentials models.Credentials) error
	CompleteRegistrationSystemAdmin(hashedEmail string) error
}

func (db *Store) CompleteRegistrationSystemAdmin(email string) error {
	// UPDATE FIELDS
	var retVal models.SystemAdmin
	if err := db.Model(&retVal).Where("email = ?", email).Update("registration_complete", true).Error; err != nil {
		return err
	}
	return nil
}

func (db *Store) RegisterSystemAdmin(credentials models.Credentials) error {
	// CREATE HASH
	hashedPassword, err := createHash(credentials.Password)
	if err != nil {
		return err
	}

	// CREATE NEW SYSTEM ADMIN
	var admin models.SystemAdmin
	admin.Credentials = credentials
	admin.Password = string(hashedPassword)

	// ADD TO DATABASE
	if err := db.Create(&admin).Error; err != nil {
		return err
	}
	return nil
}

func createHash(hashObject string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(hashObject), bcrypt.DefaultCost)
}
