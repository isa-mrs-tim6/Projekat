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
	Disabled      bool
	AirplaneID    uint
	ReservationID uint
	QuickReserve  bool
}

type AirplaneDAO struct {
	Name      string
	RowWidth  uint
	First     uint
	Business  uint
	Economy   uint
}

type Airplane struct {
	gorm.Model
	Name      string
	IsCopy    bool
	RowWidth  uint
	Seats     []Seat `gorm:"foreignkey:AirplaneID"`
	AirlineID uint
}
type Destination struct {
	gorm.Model
	AirlineID uint
	Name      string
	Coordinate
}

type PriceList struct {
	PriceECONOMY               float64
	PriceBUSINESS              float64
	PriceFIRSTCLASS            float64
	SmallSuitcase              float64
	BigSuitcase                float64
	QuickReservationPriceScale float64
}

type Layovers struct {
	gorm.Model
	Address
	FlightID uint
}

type FeatureAirline struct {
	gorm.Model
	Price       float64
	Description string
	Name        string
	Icon        string
	AirlineID   uint
}

type SeatQuickReservation struct {
	gorm.Model
	SeatNumber                string
	SeatID                    string
	FlightQuickReserveSeatsID uint
}

type FlightQuickReserveSeats struct {
	gorm.Model
	Seat          []SeatQuickReservation `gorm:"foreignkey:FlightQuickReserveSeatsID"`
	FlightID      uint
	ReservationID uint
}

type Flight struct {
	gorm.Model
	Origin        *Destination
	OriginID      uint
	Destination   *Destination
	DestinationID uint
	Departure     time.Time
	Landing       time.Time
	Duration      time.Duration
	Distance      uint
	PriceList
	Layovers          []Layovers                `gorm:"foreignkey:FlightID"`
	QuickReserveSeats []FlightQuickReserveSeats `gorm:"foreignkey:FlightID"`
	Airplane          Airplane                  `gorm:"foreignkey:AirplaneID"`
	AirplaneID        uint
	Airline           *Airline `gorm:"foreignkey:AirlineID"`
	AirlineID         uint
}

type FlightRating struct {
	FlightID      uint
	ReservationID uint
	Rating        int
}

type FlightRatingDAO struct {
	Flight Flight
	Rating int
}

type LayoverDto struct {
	Address string
}

type FlightDto struct {
	FlightID        string
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
	SmallSuitcase   string
	BigSuitcase     string
	Airplane        string
	AirplaneObject  Airplane
}

type AirlineProfile struct {
	Name string
	Address
	Promo   string
	Picture string `gorm:"default:'airlinePlaceholder.png'"`
}

type AirlineProfileDTO struct {
	ID uint
	AirlineProfile
}

type Airline struct {
	gorm.Model
	AirlineProfile
	Flights         []Flight         `gorm:"foreignkey:AirlineID"`
	Airplanes       []Airplane       `gorm:"foreignkey:AirlineID"`
	Admins          []*AirlineAdmin  `gorm:"foreignkey:AirlineID"`
	Destinations    []Destination    `gorm:"foreignkey:AirlineID"`
	AirlineFeatures []FeatureAirline `gorm:"foreignkey:AirlineID"`
}

type FlightReservationParams struct {
	Features []FeatureAirline
	BigSuitcase int
	SmallSuitcase int
	Seats          []Seat
	Users          []UserReserveParams
	IsQuickReserve bool
}
