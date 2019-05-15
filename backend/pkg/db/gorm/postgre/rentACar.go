package postgre

import (
	"errors"
	"github.com/isa-mrs-tim6/Projekat/pkg/models"
	"math"
	"strconv"
	"time"
)

type RentACarDBInterface interface {
	GetRentACarCompanies() ([]models.RentACarCompany, error)
	GetRentACarCompanyProfile(id uint) (models.RentACarCompanyProfile, error)
	UpdateRentACarCompanyProfile(id uint, newProfile models.RentACarCompanyProfile)
	FindVehicles(id uint, params models.FindVehicleParams) ([]models.Vehicle, error)
}

func (db *Store) GetRentACarCompanies() ([]models.RentACarCompany, error) {
	var retVal []models.RentACarCompany
	if err := db.Set("gorm:auto_preload", true).Find(&retVal).Error; err != nil {
		return retVal, err
	}
	return retVal, nil
}

func (db *Store) CreateRentACarCompany(rac *models.RentACarCompany) error {
	if err := db.Create(&rac).Error; err != nil {
		return err
	}
	return nil
}

func (db *Store) GetRACReservations(id uint) ([]models.RentACarReservation, error) {
	var retVal []models.RentACarReservation
	if err := db.Set("gorm:auto_preload", true).Where("company_id = ?", id).Find(&retVal).Error; err != nil {
		return retVal, err
	}
	return retVal, nil
}

func (db *Store) GetVehicleRatings(id uint) ([]models.VehicleRatingsDAO, error) {
	var retVal []models.VehicleRatingsDAO
	var vehicles []models.Vehicle
	if err := db.Set("gorm:auto_preload", true).Where("rent_a_car_company_id = ?", id).Find(&vehicles).Error; err != nil {
		return nil, err
	}

	for _, v := range vehicles {
		var ratings []models.VehicleRating
		if err := db.Where("vehicle_id = ?", v.ID).Find(&ratings).Error; err != nil {
			return nil, err
		}
		sumRating := 0
		lenRating := 1
		if len(ratings) > 0 {
			lenRating = len(ratings)
		}

		for _, rating := range ratings {
			sumRating += rating.Rating
		}
		retVal = append(retVal, models.VehicleRatingsDAO{Vehicle: v, Rating: sumRating / lenRating})
	}

	return retVal, nil
}

func (db *Store) GetRentACarCompanyProfile(id uint) (models.RentACarCompanyProfile, error) {
	var retVal models.RentACarCompany
	if err := db.First(&retVal, id).Error; err != nil {
		return retVal.RentACarCompanyProfile, err
	}
	return retVal.RentACarCompanyProfile, nil
}

func (db *Store) UpdateRentACarCompanyProfile(id uint, newProfile models.RentACarCompanyProfile) error {
	var retVal models.RentACarCompany
	if err := db.First(&retVal, id).Error; err != nil {
		return err
	}
	retVal.RentACarCompanyProfile = newProfile

	if newProfile.Promo == "" ||
		newProfile.Name == "" ||
		newProfile.Address.Address == "" {
		return errors.New("empty parameters")
	}

	if err := db.Save(&retVal).Error; err != nil {
		return err
	}
	return nil
}

func (db *Store) FindVehicles(params models.FindVehicleParams) ([]models.VehicleRatingsDAO, error) {
	var retVal []models.Vehicle
	var vehicles []models.VehicleRatingsDAO

	start, _ := strconv.ParseInt(params.StartDate, 10, 64)
	end, _ := strconv.ParseInt(params.EndDate, 10, 64)

	startDate := time.Unix(0, start*int64(time.Millisecond))
	endDate := time.Unix(0, end*int64(time.Millisecond))

	if err := db.Set("gorm:auto_preload", true).
		Where("rent_a_car_company_id = ?", params.ID).
		Where("type ILIKE ?", "%"+params.Type+"%").
		Where("capacity >= ?", params.Capacity).
		Where("price_per_day BETWEEN ? and ?", params.PriceLow, params.PriceHigh).
		Find(&retVal).Error; err != nil {
		return nil, err
	}

	for _, vehicle := range retVal {
		taken := false
		for _, res := range vehicle.Reservations {
			if !(res.Occupation.Beginning.After(endDate) ||
				res.Occupation.End.Before(startDate)) {
				taken = true
				break
			}
		}
		if !taken {
			var ratings []models.VehicleRating
			if err := db.Where("vehicle_id = ?", vehicle.ID).Find(&ratings).Error; err != nil {
				return nil, err
			}
			sumRating := 0
			lenRating := 1
			if len(ratings) > 0 {
				lenRating = len(ratings)
			}

			for _, rating := range ratings {
				sumRating += rating.Rating
			}
			vehicle.PricePerDay *= math.Ceil(endDate.Sub(startDate).Hours() / 24.0)
			vehicles = append(vehicles, models.VehicleRatingsDAO{Vehicle: vehicle, Rating: sumRating / lenRating})
		}
	}

	return vehicles, nil
}

func (db *Store) GetCompanyVehicles(id uint) ([]models.Vehicle, error) {
	var retVal []models.Vehicle
	if err := db.Set("gorm:auto_preload", true).Where("rent_a_car_company_id = ?", id).Find(&retVal).Error; err != nil {
		return retVal, err
	}
	return retVal, nil
}

func (db *Store) UpdateVehicle(id uint, newVehicle models.Vehicle) error {
	var retVal models.Vehicle
	var res []models.RentACarReservation
	if err := db.Set("gorm:auto_preload", true).Find(&res).Error; err != nil {
		return err
	}
	if err := db.First(&retVal, id).Error; err != nil {
		return err
	}
	for _, r := range res {
		if r.ID == id && r.Occupation.End.After(time.Now()) {
			return errors.New("Vehicle is reserved and cannot be updated")
		}
	}

	retVal.Name = newVehicle.Name
	retVal.Capacity = newVehicle.Capacity
	retVal.Type = newVehicle.Type
	retVal.PricePerDay = newVehicle.PricePerDay
	retVal.Discount = newVehicle.Discount

	if err := db.Save(&retVal).Error; err != nil {
		return err
	}
	return nil
}

func (db *Store) DeleteVehicle(id uint) error {
	var retVal models.Vehicle
	var res []models.RentACarReservation
	if err := db.Set("gorm:auto_preload", true).Find(&res).Error; err != nil {
		return err
	}
	if err := db.Set("gorm:auto_preload", true).First(&retVal, id).Error; err != nil {
		return err
	}

	for _, r := range res {
		if r.ID == id && r.Occupation.End.After(time.Now()) {
			return errors.New("Vehicle is reserved and cannot be deleted")
		}
	}

	if err := db.Delete(&retVal).Error; err != nil {
		return err
	}
	return nil
}

func (db *Store) AddVehicle(vehicle models.Vehicle) error {
	var retVal models.Vehicle

	retVal = vehicle

	if err := db.Create(&retVal).Error; err != nil {
		return err
	}
	return nil
}

func (db *Store) GetCompanyLocations(id uint) ([]models.Location, error) {
	var retVal []models.Location
	if err := db.Set("gorm:auto_preload", true).Where("rent_a_car_company_id = ?", id).Find(&retVal).Error; err != nil {
		return retVal, err
	}
	return retVal, nil
}

func (db *Store) UpdateLocation(id uint, newOffice models.Location) error {
	var retVal models.Location
	if err := db.First(&retVal, id).Error; err != nil {
		return err
	}

	retVal.Address = newOffice.Address
	retVal.Longitude = newOffice.Longitude
	retVal.Latitude = newOffice.Latitude

	if err := db.Save(&retVal).Error; err != nil {
		return err
	}
	return nil
}

func (db *Store) DeleteLocation(id uint) error {
	var retVal models.Location
	if err := db.First(&retVal, id).Error; err != nil {
		return err
	}
	if err := db.Delete(&retVal).Error; err != nil {
		return err
	}
	return nil
}

func (db *Store) AddLocation(location models.Location) error {
	var retVal models.Location

	retVal = location

	if err := db.Create(&retVal).Error; err != nil {
		return err
	}
	return nil
}
