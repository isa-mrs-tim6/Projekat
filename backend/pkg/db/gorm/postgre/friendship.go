package postgre

import "github.com/isa-mrs-tim6/Projekat/pkg/models"

func (db *Store) GetFriends(id uint) ([]models.User, error) {
	var retValFirstHalf []int
	var retValSecondHalf []int
	var fullList []int
	var retVal []models.User
	if err := db.Table("friendships").
		Where("user2_id = ? AND status LIKE 'ACCEPTED'", id).Pluck("user1_id", &retValFirstHalf).Error; err != nil {
		return retVal, err
	}
	if err := db.Table("friendships").
		Where("user1_id = ? AND status LIKE 'ACCEPTED'", id).Pluck("user2_id", &retValSecondHalf).Error; err != nil {
		return retVal, err
	}
	fullList = append(retValFirstHalf, retValSecondHalf...)

	if err := db.Where(fullList).Find(&retVal).Error; err != nil {
		return retVal, err
	}

	return retVal, nil
}
