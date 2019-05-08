package postgre

import "github.com/isa-mrs-tim6/Projekat/pkg/models"

func (db *Store) GetAirlines() ([]models.Airline, error) {
	var retVal []models.Airline
	if err := db.Set("gorm:auto_preload", true).Find(&retVal).Error; err != nil {
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
	if err := db.Set("gorm:auto_preload", true).Where("airline_id = ?", id).Find(&flights).Error; err != nil {
		return nil, err
	}

	for _, flight := range flights {
		var ratings []models.FlightReservation
		if err := db.Where("flight_id = ?", flight.ID).Find(&ratings).Error; err != nil {
			return nil, err
		}
		sumRating := 0
		lenRating := 1
		if len(ratings) > 0 {
			lenRating = len(ratings)
		}

		for _, rating := range ratings {
			sumRating += int(rating.Rating)
		}
		retVal = append(retVal, models.FlightRatingDAO{Flight: flight, Rating: sumRating / lenRating})
	}

	return retVal, nil
}

func (db *Store) GetAirlineReservations(id uint) ([]models.FlightReservation, error) {
	var retVal []models.FlightReservation
	if err := db.Set("gorm:auto_preload", true).
		Joins("JOIN flights on flight_reservations.flight_id=flights.id").
		Where("airline_id = ?", id).Find(&retVal).Error; err != nil {
		return retVal, err
	}
	return retVal, nil
}