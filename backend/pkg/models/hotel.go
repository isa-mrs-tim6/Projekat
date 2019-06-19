package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type RoomQuickReserveDays struct {
	gorm.Model
	Start  time.Time
	End    time.Time
	RoomID uint
}

type RoomQuickReserveDaysDAO struct {
	ID     uint
	Start  string
	End    string
	RoomID uint
}

type Room struct {
	gorm.Model
	Number           int
	Price            float64
	Capacity         uint
	HotelID          uint
	QuickReserveDays []RoomQuickReserveDays `gorm:"foreignkey:RoomID"`
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
	Picture     string `gorm:"default:'hotelPlaceholder.png'"`
}

type Hotel struct {
	gorm.Model
	HotelProfile
	Rooms    []Room                   `gorm:"foreignkey:HotelID"`
	Features []Feature                `gorm:"foreignkey:HotelID"`
	Rewards  []HotelReservationReward `gorm:"foreignkey:HotelID"`
}

type HotelReservationReward struct {
	gorm.Model
	Features    []*Feature `gorm:"many2many:reward_features;"`
	Name        string
	Description string
	PriceScale  float64
	HotelID     uint
}

type HotelReservationParamsDTO struct {
	From           string
	To             string
	Rooms          []Room
	Features       []Feature
	IsQuickReserve bool
}

type HotelReservationParams struct {
	From           time.Time
	To             time.Time
	Rooms          []Room
	Features       []*Feature
	IsQuickReserve bool
}

type HotelSearchResults struct {
	Hotel  Hotel
	Rating float64
}

type RoomSearchResults struct {
	Room   Room
	Rating float64
	Price  float64
}
