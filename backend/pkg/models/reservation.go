package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Occupation struct {
	Beginning time.Time
	End       time.Time
}

type ReservationReward struct {
	RequiredNumber uint
	PriceScale     float64
}

type ReservationDAO struct {
	Master Reservation
	Slaves []Reservation
}

type ReservationGraphData struct {
	Id        uint
	Departure time.Time
	Price     float64
}

type Reservation struct {
	gorm.Model
	Passenger
	MasterRef             uint
	ReservationFlight     FlightReservation `gorm:"foreignkey:ReservationFlightID"`
	ReservationFlightID   uint
	ReservationRentACar   RentACarReservation `gorm:"foreignkey:ReservationRentACarID"`
	ReservationRentACarID uint
	ReservationHotel      HotelReservation `gorm:"foreignkey:ReservationHotelID"`
	ReservationHotelID    uint
}

type FlightReservation struct {
	gorm.Model
	Flight         *Flight `gorm:"foreignkey:FlightID"`
	FlightID       uint
	Seat           *Seat `gorm:"foreignkey:ReservationID"`
	Price          float64
	Features       []*FeatureAirline `gorm:"many2many:flight_reservation_feature;"`
	CompanyRating  uint
	FlightRating   uint
	IsQuickReserve bool
}

type RentACarReservation struct {
	gorm.Model
	Occupation
	Location        string
	RentACarCompany RentACarCompany `gorm:"foreignkey:CompanyID"`
	CompanyID       uint
	Vehicle         Vehicle `gorm:"foreignkey:VehicleID"`
	VehicleID       uint
	Price           float64
	CompanyRating   uint
	VehicleRating   uint
	IsQuickReserve  bool
}

type HotelReservation struct {
	gorm.Model
	Occupation
	Hotel          Hotel `gorm:"foreignkey:HotelID"`
	HotelID        uint
	Rooms          []Room       `gorm:"many2many:room_reservations;"`
	Ratings        []RoomRating `gorm:"foreignkey:ReservationID"`
	Features       []*Feature   `gorm:"many2many:room_reservation_feature;"`
	HotelRating    uint
	Price          float64
	CompanyRating  uint
	IsQuickReserve bool
}
