package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"time"
)

type Seat struct {
	gorm.Model
	Number        uint
	Class         string
	AirplaneID    uint
	ReservationID uint
	QuickReserve  bool
}

type Airplane struct {
	gorm.Model
	Name      string
	Seats     []Seat `gorm:"foreignkey:AirplaneID"`
	AirlineID uint
}
type Destination struct {
	gorm.Model
	AirlineID uint
	Name      string
	Coordinate
}

type Layovers struct {
	gorm.Model
	Address
	FlightID uint
}

type Flight struct {
	gorm.Model
	Origin                     *Destination
	OriginID                   uint
	Destination                *Destination
	DestinationID              uint
	Departure                  time.Time
	Landing                    time.Time
	Duration                   time.Duration
	Distance                   uint
	Layovers                   []Layovers `gorm:"foreignkey:FlightID"`
	PriceECONOMY               float64
	PriceBUSINESS              float64
	PriceFIRSTCLASS            float64
	QuickReservationPriceScale float64
	Airplane                   Airplane `gorm:"foreignkey:AirplaneID"`
	AirplaneID                 uint
	AirlineID                  uint
	Reservations               []FlightReservation `gorm:"foreignkey:FlightID"`
}

type LayoverDto struct {
	Address string
}

type FlightDto struct {
	Origin          string
	Destination     string
	Departure       string
	Landing         string
	Duration        string
	Distance        string
	Layovers        []LayoverDto
	PriceECONOMY    string
	PriceBUSINESS   string
	PriceFIRSTCLASS string
	Airplane        string
}

type AirlineProfile struct {
	Name string
	Address
	Promo string
}

type Airline struct {
	gorm.Model
	AirlineProfile
	Flights      []Flight        `gorm:"foreignkey:AirlineID"`
	Airplanes    []Airplane      `gorm:"foreignkey:AirlineID"`
	Admins       []*AirlineAdmin `gorm:"foreignkey:AirlineID"`
	Destinations []Destination   `gorm:"foreignkey:AirlineID"`
}
