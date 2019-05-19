package postgre

import (
	"github.com/isa-mrs-tim6/Projekat/pkg/models"
	"math"
	"strings"
)

func (db *Store) OneWaySearch(query models.OneWayQuery) ([]models.Flight, error) {
	var flights []models.Flight
	if err := db.Preload("Origin").Preload("Destination").Preload("Airline").
		Joins("JOIN destinations fromDest on fromDest.id=flights.origin_id").
		Where("LOWER(fromDest.name) LIKE ?", "%"+query.From+"%").
		Joins("JOIN destinations toDest on toDest.id=flights.destination_id").
		Where("LOWER(toDest.name) LIKE ? and date(flights.departure)=?", "%"+query.To+"%", query.Date).
		Joins("JOIN airplanes on airplanes.id=flights.airplane_id").
		Joins("JOIN seats on flights.airplane_id=seats.airplane_id").
		Where("seats.class=? and seats.reservation_id = 0", query.SeatClass).
		Group("flights.id").
		Having("count(flights.id)>=?", query.Passengers).
		Find(&flights).Error; err != nil {
		return nil, err
	}
	return flights, nil
}

func (db *Store) HotelSearch(query models.HotelQuery) ([]models.Hotel, error) {
	var hotels []models.Hotel

	query.Name = strings.ToLower(query.Name)
	query.Address = strings.ToLower(query.Address)
	if query.RoomCapacityDownLimit == 0 && query.RoomCapacityUpLimit == 0 {
		query.RoomCapacityUpLimit = math.MaxUint8
	}
	if query.RoomPriceDownLimit == 0 && query.RoomPriceUpLimit == 0 {
		query.RoomPriceUpLimit = math.MaxFloat64
	}

	if err := db.Joins("JOIN rooms rooms on rooms.hotel_id=hotels.id").
		Where("rooms.price BETWEEN ? AND ? and rooms.capacity BETWEEN ? AND ?",
			query.RoomPriceDownLimit, query.RoomPriceUpLimit, query.RoomCapacityDownLimit, query.RoomCapacityUpLimit).
		Where("LOWER(hotels.name) LIKE ? AND LOWER(hotels.address) LIKE ?", "%"+query.Name+"%", "%"+query.Address+"%").
		Group("hotels.id").
		Find(&hotels).Error; err != nil {
		return nil, err
	}

	return hotels, nil
}
