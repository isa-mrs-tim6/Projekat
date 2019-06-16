package postgre

import (
	"github.com/isa-mrs-tim6/Projekat/pkg/models"
)

func (db *Store) GetUsers() ([]models.User, error) {
	var retVal []models.User
	if err := db.Set("gorm:auto_preload", true).Find(&retVal).Error; err != nil {
		return retVal, err
	}
	return retVal, nil
}

func (db *Store) UpdateUser(oldEmail string, params models.ProfileParams) error {
	var retVal models.User

	if err := db.Where("email = ?", oldEmail).First(&retVal).Error; err != nil {
		return err
	}

	hashedPassword, err := createHash(params.Profile.Password)
	if err != nil {
		return err
	}
	params.Profile.Password = string(hashedPassword)

	retVal.Profile = params.Profile

	if err := db.Save(&retVal).Error; err != nil {
		return err
	}
	return nil
}

func (db *Store) GetUser(email string) (models.User, error) {
	var retVal models.User
	if err := db.Where("email = ?", email).First(&retVal).Error; err != nil {
		return retVal, err
	}
	return retVal, nil
}

func (db *Store) GetUserProfile(email string) (models.Profile, error) {
	var retVal models.User
	if err := db.Where("email = ?", email).First(&retVal).Error; err != nil {
		return retVal.Profile, err
	}
	return retVal.Profile, nil
}

func (db *Store) GetAirlineAdmin(email string) (models.AirlineAdmin, error) {
	var retVal models.AirlineAdmin
	if err := db.Where("email = ?", email).First(&retVal).Error; err != nil {
		return retVal, err
	}
	return retVal, nil
}

func (db *Store) UpdateAirlineAdmin(email string, newProfile models.Profile) error {
	var retVal models.AirlineAdmin
	if err := db.First(&retVal, "email = ?", email).Error; err != nil {
		return err
	}
	hashedPassword, err := createHash(newProfile.Credentials.Password)
	if err != nil {
		return err
	}
	newProfile.Password = string(hashedPassword)

	retVal.Profile = newProfile
	retVal.FirstPassChanged = true
	if err := db.Save(&retVal).Error; err != nil {
		return err
	}
	return nil
}

func (db *Store) GetRACAdmin(email string) (models.RentACarAdmin, error) {
	var retVal models.RentACarAdmin
	if err := db.Where("email = ?", email).First(&retVal).Error; err != nil {
		return retVal, err
	}
	return retVal, nil
}

func (db *Store) UpdateRACAdmin(email string, newProfile models.Profile) error {
	var retVal models.RentACarAdmin
	if err := db.First(&retVal, "email = ?", email).Error; err != nil {
		return err
	}
	hashedPassword, err := createHash(newProfile.Credentials.Password)
	if err != nil {
		return err
	}
	newProfile.Password = string(hashedPassword)

	retVal.Profile = newProfile
	retVal.FirstPassChanged = true
	if err := db.Save(&retVal).Error; err != nil {
		return err
	}
	return nil
}

func (db *Store) GetHotelAdmin(email string) (models.HotelAdmin, error) {
	var retVal models.HotelAdmin
	if err := db.Where("email = ?", email).First(&retVal).Error; err != nil {
		return retVal, err
	}
	return retVal, nil
}

func (db *Store) UpdateHotelAdmin(email string, newProfile models.Profile) error {
	var retVal models.HotelAdmin
	if err := db.First(&retVal, "email = ?", email).Error; err != nil {
		return err
	}
	hashedPassword, err := createHash(newProfile.Credentials.Password)
	if err != nil {
		return err
	}
	newProfile.Password = string(hashedPassword)

	retVal.Profile = newProfile
	retVal.FirstPassChanged = true
	if err := db.Save(&retVal).Error; err != nil {
		return err
	}
	return nil
}

func (db *Store) GetSystemAdmin(email string) (models.SystemAdmin, error) {
	var retVal models.SystemAdmin
	if err := db.Where("email = ?", email).First(&retVal).Error; err != nil {
		return retVal, err
	}
	return retVal, nil
}

func (db *Store) UpdateSystemAdmin(email string, newProfile models.Profile) error {
	var retVal models.SystemAdmin
	if err := db.First(&retVal, "email = ?", email).Error; err != nil {
		return err
	}

	hashedPassword, err := createHash(newProfile.Credentials.Password)
	if err != nil {
		return err
	}
	newProfile.Password = string(hashedPassword)

	retVal.Profile = newProfile
	retVal.FirstPassChanged = true
	if err := db.Save(&retVal).Error; err != nil {
		return err
	}
	return nil
}

func (db *Store) Rate(rating models.ResRatingDAO) error {
	switch rating.ResType {
	case "flight":
		var fRes models.FlightReservation
		if err := db.First(&fRes, rating.ResID).Error; err != nil {
			return err
		}
		fRes.FlightRating = rating.Rating
		if err := db.Save(&fRes).Error; err != nil {
			return err
		}
		break
	case "airline":
		var fRes models.FlightReservation
		if err := db.First(&fRes, rating.ResID).Error; err != nil {
			return err
		}
		fRes.CompanyRating = rating.Rating
		if err := db.Save(&fRes).Error; err != nil {
			return err
		}
		break
	case "hotel":
		var hRes models.HotelReservation
		if err := db.First(&hRes, rating.ResID).Error; err != nil {
			return err
		}
		hRes.HotelRating = rating.Rating
		if err := db.Save(&hRes).Error; err != nil {
			return err
		}
		break
	case "room":
		var ratings []models.RoomRating
		db.Table("room_ratings").Find(&ratings)
		found := false
		for _, r := range ratings {
			if r.ReservationID == rating.ResID && r.RoomID == rating.RoomID {
				r.Rating = int(rating.Rating)
				db.Table("room_ratings").
					Where("room_id = ?", rating.RoomID).
					Where("reservation_id = ?", rating.ResID).
					Update("rating", int(rating.Rating))
				found = true
				break
			}
		}
		if !found {
			var rating = models.RoomRating{
				Rating:        int(rating.Rating),
				RoomID:        rating.RoomID,
				ReservationID: rating.ResID,
			}
			db.Create(&rating)
		}
		break
	case "vehicle":
		var rRes models.RentACarReservation
		if err := db.First(&rRes, rating.ResID).Error; err != nil {
			return err
		}
		rRes.VehicleRating = rating.Rating
		if err := db.Save(&rRes).Error; err != nil {
			return err
		}
		break
	case "rac":
		var rRes models.RentACarReservation
		if err := db.First(&rRes, rating.ResID).Error; err != nil {
			return err
		}
		rRes.CompanyRating = rating.Rating
		if err := db.Save(&rRes).Error; err != nil {
			return err
		}
		break
	}
	return nil
}
