package models

import "github.com/jinzhu/gorm"

type Credentials struct {
	Email    string `gorm: "type:varchar(100); unique_index"`
	Password string
}

type UserInfo struct {
	Name    string
	Surname string
}

type User struct {
	gorm.Model
	Credentials
	UserInfo
	Reservations []*Reservation `gorm:"many2many:user_reservations;"`
}

type AirlineAdmin struct {
	gorm.Model
	Credentials
	UserInfo
	AirlineID uint
}

type HotelAdmin struct {
	gorm.Model
	Credentials
	UserInfo
	HotelID uint
}

type RentACarAdmin struct {
	gorm.Model
	Credentials
	UserInfo
	RentACarCompanyID uint
}

type SystemAdmin struct {
	gorm.Model
	Credentials
	UserInfo
}
