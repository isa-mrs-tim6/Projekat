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

type ReservationGraphData struct {
	Id        uint
	Departure time.Time
	Price     float64
}

func (r *Reservation) BeforeCreate(scope *gorm.Scope) (err error) {
	scope.DB().Where("is_expiring = true and expire_time < ?", time.Now()).Delete(Reservation{})

	var ids []uint

	scope.DB().Table("reservations").
		Joins("JOIN flight_reservations on reservations.reservation_flight_id = flight_reservations.id").
		Joins("JOIN flights on flight_reservations.flight_id = flights.id").
		Where("reservations.is_expiring = true AND flights.departure < ?", time.Now().Add(time.Hour*2)).
		Pluck("reservations.id", &ids)

	if ids != nil {
		scope.DB().Table("reservations").
			Where("id in ?", ids).
			Delete(Reservation{})
	}

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
