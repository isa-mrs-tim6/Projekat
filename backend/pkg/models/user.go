package models

import "github.com/jinzhu/gorm"

type Credentials struct {
	Email    string `gorm:"type:varchar(100); unique_index"`
	Password string
}

type UserInfo struct {
	Name    string
	Surname string
}

type Profile struct {
	Credentials
	UserInfo
	Address string
	Phone   string
}

type User struct {
	gorm.Model
	Profile
	RegistrationComplete bool
	Reservations         []*Reservation `gorm:"many2many:user_reservations;"`
}

type AirlineAdmin struct {
	gorm.Model
	Profile
	RegistrationComplete bool
	FirstPassChanged     bool
	AirlineID            uint
}

type HotelAdmin struct {
	gorm.Model
	Profile
	RegistrationComplete bool
	FirstPassChanged     bool
	HotelID              uint
}

type RentACarAdmin struct {
	gorm.Model
	Profile
	RegistrationComplete bool
	FirstPassChanged     bool
	RentACarCompanyID    uint
}

type SystemAdmin struct {
	gorm.Model
	Profile
	RegistrationComplete bool
	FirstPassChanged     bool
}
