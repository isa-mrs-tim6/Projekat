package models

type Coordinate struct {
	Latitude  float64
	Longitude float64
}

type Address struct {
	Address string
	Coordinate
}
