package postgre

import (
	"errors"
	"github.com/isa-mrs-tim6/Projekat/pkg/models"
	"golang.org/x/crypto/bcrypt"
)

func (db *Store) Login(credentials models.Credentials) error {

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
