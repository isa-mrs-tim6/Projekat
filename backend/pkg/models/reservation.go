package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Occupation struct {
	Beginning time.Time
	End       time.Time
}

type ReservationReward struct {
	RequiredNumber uint
	PriceScale     float64
}

type ReservationScaleDAO struct {
	Count uint
	Scale float64
}

type ReservationDAO struct {
	Master    Reservation
	InvitedBy User
	Slaves    []Reservation
}

type QuickFlightReservationDTO struct {
	SeatID   uint
	FlightID uint
	Discount uint
}

type QuickFlightReservationGDTO struct {
	Number        uint
	Price         float64
	ReservationID uint
	ID            uint
}

type QuickFlightReservationGDTOV2 struct {
	Class           string
	Price           float64
	ID              uint
	OriginName      string
	DestName        string
	Departure       time.Time
	FlightID        uint
	PriceECONOMY    string
	PriceBUSINESS   string
	PriceFIRSTCLASS string
}

type ReservationGraphData struct {
	Id        uint
	Departure time.Time
	Price     float64
}

func (r *Reservation) BeforeCreate(scope *gorm.Scope) (err error) {
	var reservations []Reservation

	scope.DB().Where("is_expiring = true and expire_time < ?", time.Now()).Find(&reservations)

	for _, res := range reservations {
		Cancel(scope, res.ID)
	}

	scope.DB().Table("reservations").
		Joins("JOIN flight_reservations on reservations.reservation_flight_id = flight_reservations.id").
		Joins("JOIN flights on flight_reservations.flight_id = flights.id").
		Where("reservations.is_expiring = true AND flights.departure < ?", time.Now().Add(time.Hour*3)).
		Find(&reservations)

	for _, res := range reservations {
		Cancel(scope, res.ID)
	}

	return
}

func Cancel(scope *gorm.Scope, resID uint) {
	var master Reservation

	scope.DB().Where("id = ?", resID).
		Preload("ReservationFlight.Flight.Origin").
		Preload("ReservationFlight.Flight.Destination").
		Preload("ReservationFlight.Flight.Layovers").
		Preload("ReservationFlight.Seat").
		Preload("ReservationFlight.Features").
		Preload("ReservationHotel.Rooms").
		Preload("ReservationHotel.Ratings").
		Preload("ReservationHotel.Features").
		Preload("ReservationHotel.Hotel").
		Preload("ReservationRentACar.RentACarCompany").
		Preload("ReservationRentACar.Vehicle").
		First(&master)
	scope.DB().Table("seats").Where("reservation_id = ?", master.ReservationFlight.Seat.ReservationID).
		Update("reservation_id", 0)

	if !master.ReservationFlight.IsQuickReserve {
		scope.DB().Where("id=?", master.ID).Delete(master.ReservationFlight)
	}
	if !master.ReservationHotel.IsQuickReserve {
		scope.DB().Where("id=?", master.ID).Delete(master.ReservationHotel)
	}
	if !master.ReservationRentACar.IsQuickReserve {
		scope.DB().Where("id=?", master.ID).Delete(master.ReservationRentACar)
	}

	scope.DB().Delete(&master)

	return

}

type Reservation struct {
	gorm.Model
	Passenger
	MasterRef             uint
	ReservationFlight     FlightReservation `gorm:"foreignkey:ReservationFlightID"`
	ReservationFlightID   uint
	ReservationRentACar   RentACarReservation `gorm:"foreignkey:ReservationRentACarID"`
	ReservationRentACarID uint
	ReservationHotel      HotelReservation `gorm:"foreignkey:ReservationHotelID"`
	ReservationHotelID    uint
	IsExpiring            bool
	ExpireTime            time.Time
}

type FlightReservation struct {
	gorm.Model
	Flight         *Flight `gorm:"foreignkey:FlightID"`
	FlightID       uint
	Seat           *Seat `gorm:"foreignkey:ReservationID"`
	Price          float64
	Features       []*FeatureAirline `gorm:"many2many:flight_reservation_feature;"`
	CompanyRating  uint
	FlightRating   uint
	IsQuickReserve bool
}

type RentACarReservation struct {
	gorm.Model
	Occupation
	Location                   string
	RentACarCompany            RentACarCompany `gorm:"foreignkey:CompanyID"`
	CompanyID                  uint
	Vehicle                    Vehicle `gorm:"foreignkey:VehicleID"`
	VehicleID                  uint
	Price                      float64
	CompanyRating              uint
	VehicleRating              uint
	IsQuickReserve             bool
	QuickReservationPriceScale float64
}

type HotelReservation struct {
	gorm.Model
	Occupation
	Hotel          Hotel `gorm:"foreignkey:HotelID"`
	HotelID        uint
	Rooms          []Room       `gorm:"many2many:room_reservations;"`
	Ratings        []RoomRating `gorm:"foreignkey:ReservationID"`
	Features       []*Feature   `gorm:"many2many:room_reservation_feature;"`
	HotelRating    uint
	Price          float64
	CompanyRating  uint
	IsQuickReserve bool
}
