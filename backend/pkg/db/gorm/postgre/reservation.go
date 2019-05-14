package postgre

import "github.com/isa-mrs-tim6/Projekat/pkg/models"

func (db *Store) GetReservationRewards() ([]models.ReservationReward, error) {
	var retVal []models.ReservationReward
	if err := db.Set("gorm:auto_preload", true).Find(&retVal).Error; err != nil {
		return retVal, err
	}
	return retVal, nil
}

func (db *Store) UpdateReservationRewards(rewards []models.ReservationReward) {
	db.Delete(models.ReservationReward{})
	for _, v := range rewards {
		db.Create(v)
	}
}

func (db *Store) GetReservationGraphData(id uint) ([]models.ReservationGraphData, error) {
	var retVal []models.ReservationGraphData
	if err := db.Table("flight_reservations").Select("flight_reservations.id, flight_reservations.price, flights.departure").
		Joins("JOIN flights ON flights.id = flight_reservations.flight_id").Where("flights.airline_id = ?",id).Scan(&retVal).
		Error; err != nil {
		return nil, err
	}
	return retVal, nil
}