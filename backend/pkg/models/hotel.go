package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Room struct {
	gorm.Model
	Number       int
	Price        float64
	Capacity     uint
	HotelID      uint
	QuickReserve bool
}

type RoomRating struct {
	RoomID        uint
	ReservationID uint
	Rating        int
}

type RoomRatingDAO struct {
	Room   Room
	Rating int
}

type Feature struct {
	gorm.Model
	Price       float64
	Description string
	Name        string
	Icon        string
	HotelID     uint
}

type HotelProfile struct {
	Name string
	Address
	Description string
}

type Hotel struct {
	gorm.Model
	HotelProfile
	Rooms    []Room    `gorm:"foreignkey:HotelID"`
	Features []Feature `gorm:"foreignkey:HotelID"`
}

type HotelReservationParamsDTO struct {
	From           string
	To             string
	Rooms          []Room
	IsQuickReserve bool
}

type HotelReservationParams struct {
	From           time.Time
	To             time.Time
	Rooms          []Room
	IsQuickReserve bool
}
