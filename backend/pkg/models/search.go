package models

import "time"

type UserQueryDto struct{
	Query string
}

type UserResultDTO struct {
	ID uint
	Name string
	Surname string
	Email string
	Status string
}

type OneWayQueryDto struct {
	From       string
	To         string
	Date       string
	Passengers string
	SeatClass  string
	Layovers   string
}

type OneWayQuery struct {
	From       string
	To         string
	Date       string
	Passengers uint
	SeatClass  string
}

type HotelQueryDTO struct {
	Name    string
	Address string
	From    string
	To      string
}

type HotelQuery struct {
	Name    string
	Address string
	From    time.Time
	To      time.Time
}

type RoomQueryDTO struct {
	Capacities []uint
	From       string
	To         string
}

type RoomQuery struct {
	HotelID    uint
	Capacities []uint
	From       time.Time
	To         time.Time
}

type RacQueryDTO struct {
	Name    string
	Address string
	From    string
	To      string
}

type RacQuery struct {
	Name    string
	Address string
	From    time.Time
	To      time.Time
}
