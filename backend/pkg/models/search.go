package models

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

type HotelQuery struct {
	Name         string
	Address      string
	RoomCapacity string
	RoomPrice    float64
}
