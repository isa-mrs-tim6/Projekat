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

type Airplane struct {
	gorm.Model
	Name      string
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
	Layovers   []Layovers `gorm:"foreignkey:FlightID"`
	Airplane   Airplane   `gorm:"foreignkey:AirplaneID"`
	Ratings    []FlightRating
	AirplaneID uint
	AirlineID  uint
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
	Promo string
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
