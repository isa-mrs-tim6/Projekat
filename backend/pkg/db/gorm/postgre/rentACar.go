package postgre

import (
	"errors"
	"github.com/isa-mrs-tim6/Projekat/pkg/models"
	"strings"
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

func (db *Store) FindVehicles(id uint, params models.FindVehicleParams) ([]models.Vehicle, error) {
	var retVal []models.Vehicle
	var company models.RentACarCompany

	if err := db.Set("gorm:auto_preload", true).First(&company, id).Error; err != nil {
		return nil, err
	}

	for _, v := range company.Vehicles {
		if strings.Contains(strings.ToLower(v.Name), strings.ToLower(params.Name)) &&
			strings.Contains(strings.ToLower(v.Type), strings.ToLower(params.Type)) &&
			v.Capacity >= params.Capacity &&
			v.PricePerDay > params.PriceLow && v.PricePerDay < params.PriceHigh &&
			!v.Discount {
			retVal = append(retVal, v)
		}
	}

	return retVal, nil
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
	if err := db.First(&retVal, id).Error; err != nil {
		return err
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
