package postgre

import (
	"github.com/isa-mrs-tim6/Projekat/pkg/models"
)

func (db *Store) GetAirlines() ([]models.Airline, error) {
	var retVal []models.Airline
	if err := db.Find(&retVal).Error; err != nil {
		return retVal, err
	}
	return retVal, nil
}

func (db *Store) CreateAirline(airline *models.Airline) error {
	if err := db.Create(&airline).Error; err != nil {
		return err
	}
	return nil
}

func (db *Store) GetAirlineProfile(id uint) (models.AirlineProfile, error) {
	var retVal models.Airline
	if err := db.First(&retVal, id).Error; err != nil {
		return retVal.AirlineProfile, err
	}
	return retVal.AirlineProfile, nil
}

func (db *Store) GetAirlineRating(id uint) (float64, error) {
	var retVal []float64
	if err := db.Table("flight_reservations").
		Joins("INNER JOIN flights on flights.id = flight_reservations.flight_id").
		Where("flight_reservations.company_rating != 0 and flights.airline_id = ?", id).
		Pluck("avg(flight_reservations.company_rating)", &retVal).Error; err != nil {
		return retVal[0], err
	}
	return retVal[0], nil
}

func (db *Store) GetAirlinesProfiles() ([]models.AirlineProfileDTO, error) {
	var airlines []models.Airline
	if err := db.Find(&airlines).Error; err != nil {
		return nil, err
	}
	var retVal = make([]models.AirlineProfileDTO, 0)
	for _, element := range airlines{
		retVal = append(retVal, models.AirlineProfileDTO{ID: element.ID, AirlineProfile: element.AirlineProfile})
	}
	return retVal, nil
}

func (db *Store) UpdateAirline(id uint, newProfile models.AirlineProfile) error {
	var retVal models.Airline
	if err := db.First(&retVal, id).Error; err != nil {
		return err
	}
	retVal.AirlineProfile = newProfile

	if err := db.Save(&retVal).Error; err != nil {
		return err
	}
	return nil
}

func (db *Store) GetFlightRatings(id uint) ([]models.FlightRatingDAO, error) {
	var retVal []models.FlightRatingDAO
	var flights []models.Flight
	if err := db.Preload("Origin").Preload("Destination").
		Where("airline_id = ?", id).Find(&flights).Error; err != nil {
		return nil, err
	}

	for _, flight := range flights {
		var reservations []models.FlightReservation
		if err := db.Where("flight_id = ?", flight.ID).Find(&reservations).Error; err != nil {
			return nil, err
		}
		sumRating := 0
		lenRating := 1
		if len(reservations) > 0 {
			lenRating = len(reservations)
		}

		for _, reservation := range reservations {
			sumRating += int(reservation.FlightRating)
		}
		retVal = append(retVal, models.FlightRatingDAO{Flight: flight, Rating: sumRating / lenRating})
	}

	return retVal, nil
}

func (db *Store) GetAirlineReservations(id uint) ([]models.FlightReservation, error) {
	var retVal []models.FlightReservation
	if err := db.Joins("JOIN flights on flight_reservations.flight_id=flights.id").
		Where("airline_id = ?", id).Find(&retVal).Error; err != nil {
		return retVal, err
	}
	return retVal, nil
}
