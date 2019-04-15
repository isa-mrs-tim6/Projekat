package postgre

import "github.com/isa-mrs-tim6/Projekat/pkg/models"

func (db *Store) GetDestinations() ([]models.Destination, error) {
	var retVal []models.Destination
	if err := db.Set("gorm:auto_preload", true).Find(&retVal).Error; err != nil {
		return retVal, err
	}
	return retVal, nil
}

func (db *Store) GetCompanysDestinations(id uint) ([]models.Destination, error) {
	var retVal []models.Destination
	if err := db.Set("gorm:auto_preload", true).Where("airline_id = ?", id).Find(&retVal).Error; err != nil {
		return retVal, err
	}
	return retVal, nil
}

func (db *Store) FindDestination(id uint, name string) (models.Destination, error) {
	var retVal models.Destination
	if err := db.Where("airline_id = ? AND name = ?", id, name).Find(&retVal).Error; err != nil {
		return retVal, err
	}
	return retVal, nil
}

func (db *Store) GetDestination(id uint) (models.Destination, error) {
	var retVal models.Destination
	if err := db.First(&retVal, id).Error; err != nil {
		return retVal, err
	}
	return retVal, nil
}

func (db *Store) UpdateDestination(id uint, newDestination models.Destination) error {
	var retVal models.Destination
	if err := db.First(&retVal, id).Error; err != nil {
		return err
	}
	retVal = newDestination

	if err := db.Save(&retVal).Error; err != nil {
		return err
	}
	return nil
}

func (db *Store) CreateDestination(destination *models.Destination) error {
	if err := db.Create(&destination).Error; err != nil {
		return err
	}
	return nil
}
