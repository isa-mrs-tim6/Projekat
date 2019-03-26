package models

import "github.com/jinzhu/gorm"

type Vehicle struct {
	gorm.Model
	Name              string
	Capacity          uint
	PricePerDay       float64
	RentACarCompanyID uint
	ReservationID     uint
}

type Location struct {
	gorm.Model
	Address
	RentACarCompanyID uint
}

type RentACarCompany struct {
	gorm.Model
	Name string
	Address
	Promo        string
	Vehicles     []Vehicle        `gorm:"foreignkey:RentACarCompanyID"`
	Locations    []Location       `gorm:"foreignkey:RentACarCompanyID"`
	Admins       []*RentACarAdmin `gorm:"foreignkey:RentACarCompanyID"`
	Reservations []Reservation    `gorm:"foreignkey:RentACarCompanyID"`
}
