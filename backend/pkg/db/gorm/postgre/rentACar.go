package postgre

import (
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
		if strings.Contains(strings.ToLower(v.Type), strings.ToLower(params.Type)) &&
			v.Capacity >= params.Capacity &&
			v.PricePerDay > params.PriceLow && v.PricePerDay < params.PriceHigh &&
			!v.Discount {
			retVal = append(retVal, v)
		}
	}

	return retVal, nil
}
