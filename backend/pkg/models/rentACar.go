package models

import (
	"github.com/jinzhu/gorm"
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
	Ratings           []VehicleRating `gorm:"foreignkey:VehicleReferer"`
}

type VehicleParams struct {
	ID          uint
	Name        string
	Capacity    uint
	Type        string
	PricePerDay float64
	Discount    bool
}

type VehicleReservationParams struct {
	VehicleID  uint
	LocationID uint
	CompanyID  uint
	Price      float64
	StartDate  string
	EndDate    string
}

type FindVehicleParams struct {
	ID        uint
	Capacity  uint
	Type      string
	PriceLow  float64
	PriceHigh float64
	StartDate string
	EndDate   string
}

type VehicleRating struct {
	VehicleID     uint
	ReservationID uint
	Rating        int
}

type VehicleRatingsDAO struct {
	Vehicle Vehicle
	Rating  int
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
