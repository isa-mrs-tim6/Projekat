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
