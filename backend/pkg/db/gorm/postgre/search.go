package postgre

import (
	"github.com/isa-mrs-tim6/Projekat/pkg/models"
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
	var retval []models.Hotel

	query.Name = strings.ToLower(query.Name)
	query.Address = strings.ToLower(query.Address)

	if err := db.Preload("Rooms").Where("LOWER(hotels.name) LIKE ? AND LOWER(hotels.address) LIKE ?", "%"+query.Name+"%", "%"+query.Address+"%").Find(&hotels).Error; err != nil {
		return nil, err
	}

	var reservations []models.HotelReservation
	if err := db.Find(&reservations).Error; err != nil {
		return nil, err
	}

	for _, hotel := range hotels {
		for _, room := range hotel.Rooms {
			available := true
			for _, reservation := range reservations {
				for _, reservedRoom := range reservation.Rooms {
					if room.ID == reservedRoom.ID && ((reservation.Beginning.After(query.From) && reservation.Beginning.Before(query.To)) || (reservation.End.After(query.From) && reservation.End.Before(query.To))) {
						available = false
						break
					}
				}
				if available {
					break
				}
			}
			if available {
				retval = append(retval, hotel)
				break
			}
		}
	}

	return retval, nil
}
