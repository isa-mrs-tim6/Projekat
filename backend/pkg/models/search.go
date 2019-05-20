package models

import "time"

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
