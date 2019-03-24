package models

import "github.com/jinzhu/gorm"

type Room struct {
	gorm.Model
	Number        int
	Price         float64
	Capacity      uint
	HotelID       uint
	ReservationID uint
}

type Feature struct {
	gorm.Model
	Price       float64
	Description string
	HotelID     uint
}

type Hotel struct {
	gorm.Model
	Name string
	Address
	Description  string
	Rooms        []Room             `gorm:"foreignkey:HotelID"`
	Features     []Feature          `gorm:"foreignkey:HotelID"`
	Admins       []*HotelAdmin      `gorm:"foreignkey:HotelID"`
	Reservations []HotelReservation `gorm:"foreignkey:HotelID"`
}
