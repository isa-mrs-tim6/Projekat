package postgre

import (
	"github.com/isa-mrs-tim6/Projekat/pkg/models"
	"strconv"
	"time"
)

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

func (db *Store) ReserveVehicle(params models.VehicleReservationParams) error {
	var reservation models.RentACarReservation
	var vehicle models.Vehicle
	var location models.Location

	reservation.CompanyID = params.CompanyID

	if err := db.Where("id=?", params.VehicleID).First(&vehicle).Error; err != nil {
		return err
	}

	if err := db.Where("id=?", params.LocationID).First(&location).Error; err != nil {
		return err
	}

	reservation.Vehicles = append(reservation.Vehicles, &vehicle)
	reservation.Price = params.Price
	reservation.Location = location.Address.Address

	start, _ := strconv.ParseInt(params.StartDate, 10, 64)
	end, _ := strconv.ParseInt(params.EndDate, 10, 64)

	startDate := time.Unix(0, start*int64(time.Millisecond))
	endDate := time.Unix(0, end*int64(time.Millisecond))

	reservation.Occupation.Beginning = startDate
	reservation.Occupation.End = endDate

	reservation.Expiring = true

	if err := db.Create(&reservation).Update("expire_time", reservation.CreatedAt.Add(time.Minute*10)).Error; err != nil {
		return err
	}
	return nil
}
