package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Vehicle struct {
	gorm.Model
	Name              string
	Capacity          uint
	Type              string
	PricePerDay       float64
	RentACarCompanyID uint
	Reservations      []*RentACarReservation `gorm:"many2many:vehicle_reservations;"`
	Discount          bool
}

type VehicleParams struct {
	ID          uint
	Name        string
	Capacity    uint
	Type        string
	PricePerDay float64
	Discount    bool
}

type FindVehicleParams struct {
	Name      string
	Capacity  uint
	Type      string
	PriceLow  float64
	PriceHigh float64
	StartDate time.Time
	EndDate   time.Time
}

type Location struct {
	gorm.Model
	Address
	RentACarCompanyID uint
}

type LocationParams struct {
	ID uint
	Address
}

type RentACarCompanyProfile struct {
	Name string
	Address
	Promo string
}

type RentACarCompany struct {
	gorm.Model
	RentACarCompanyProfile
	Vehicles  []Vehicle        `gorm:"foreignkey:RentACarCompanyID"`
	Locations []Location       `gorm:"foreignkey:RentACarCompanyID"`
	Admins    []*RentACarAdmin `gorm:"foreignkey:RentACarCompanyID"`
}
