package models

type OneWayQueryDto struct {
	From      string
	To  string
	Date      string
	Passengers  string
	SeatClass string
	Layovers string
}

type OneWayQuery struct{
	From string
	To string
	Date string
	Passengers uint
	SeatClass string
}