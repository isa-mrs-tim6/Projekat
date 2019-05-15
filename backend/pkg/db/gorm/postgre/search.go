package postgre

import (
	"github.com/isa-mrs-tim6/Projekat/pkg/models"
)

func (db *Store) OneWaySearch(query models.OneWayQuery) ([]models.Flight, error) {
	var flights []models.Flight
	if err := db.Set("gorm:auto_preload", true).Joins("JOIN destinations fromDest on fromDest.id=flights.origin_id").
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
	return hotels, nil
}
