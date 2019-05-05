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

type Reservation struct {
	gorm.Model
	Holders               []*User           `gorm:"many2many:user_reservations;"`
	ReservationFlight     FlightReservation `gorm:"foreignkey:ReservationFlightID"`
	ReservationFlightID   uint
	ReservationRentACar   RentACarReservation `gorm:"foreignkey:ReservationRentACarID"`
	ReservationRentACarID uint
	ReservationHotel      HotelReservation `gorm:"foreignkey:ReservationHotelID"`
	ReservationHotelID    uint
}

type FlightReservation struct {
	gorm.Model
	FlightID uint
	Seats    []Seat `gorm:"foreignkey:ReservationID"`
	Price    float64
	Rating   uint
}

type RentACarReservation struct {
	gorm.Model
	Occupation
	Location string
	Vehicles []*Vehicle `gorm:"many2many:vehicle_reservations;"`
	Price    float64
	Rating   uint
}

type HotelReservation struct {
	gorm.Model
	Occupation
	HotelID     uint
	Rooms       []Room       `gorm:"many2many:room_reservations;"`
	Ratings     []RoomRating `gorm:"foreignkey:ReservationID"`
	HotelRating uint
	Price       float64
}
