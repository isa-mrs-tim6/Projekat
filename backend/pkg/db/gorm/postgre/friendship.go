package postgre

import "github.com/isa-mrs-tim6/Projekat/pkg/models"

func (db *Store) GetFriends(id uint) ([]models.User, error) {
	var retValFirstHalf []int
	var retValSecondHalf []int
	var fullList []int
	var retVal []models.User
	if err := db.Table("friendships").
		Where("user2_id = ? AND status LIKE 'ACCEPTED' AND deleted_at IS NULL", id).Pluck("user1_id", &retValFirstHalf).Error; err != nil {
		return retVal, err
	}
	if err := db.Table("friendships").
		Where("user1_id = ? AND status LIKE 'ACCEPTED' AND deleted_at IS NULL", id).Pluck("user2_id", &retValSecondHalf).Error; err != nil {
		return retVal, err
	}
	fullList = append(retValFirstHalf, retValSecondHalf...)

	if err := db.Where(fullList).Find(&retVal).Error; err != nil {
		return retVal, err
	}

	return retVal, nil
}

func (db *Store) GetPendingRequests(id uint) ([]models.User, error) {
	var fullList []int
	var retVal []models.User
	if err := db.Table("friendships").
		Where("user2_id = ? AND status LIKE 'PENDING' AND deleted_at IS NULL", id).Pluck("user1_id", &fullList).Error; err != nil {
		return retVal, err
	}
	if err := db.Where(fullList).Find(&retVal).Error; err != nil {
		return retVal, err
	}

	return retVal, nil
}

func (db *Store) UpdateFriendRequest(friendship models.FriendshipDto) error{
	if friendship.Status == true{
		if err := db.Table("friendships").Where("user1_id = ? AND user2_id = ? AND status LIKE 'PENDING' OR user1_id = ? AND user2_id = ? AND status LIKE 'PENDING'",friendship.User1ID, friendship.User2ID, friendship.User2ID, friendship.User1ID).
			Update(models.Friendship{Status:"ACCEPTED"}).Error; err != nil{
				return err
		}
	}else{
		if err := db.Table("friendships").Where("user1_id = ? AND user2_id = ? AND status LIKE 'PENDING' OR user1_id = ? AND user2_id = ? AND status LIKE 'PENDING'",friendship.User1ID, friendship.User2ID, friendship.User2ID, friendship.User1ID).
			Update(models.Friendship{Status:"DENIED"}).Error; err != nil{
			return err
		}
	}
	return nil
}

func (db *Store) CreateRequest(friendship models.Friendship) error {
	if err := db.Create(&friendship).Error; err!= nil{
		return err
	}
	return nil
}

func (db *Store) RemoveAFriend(friendship models.FriendshipDto) error {
	if err := db.Where("user1_id = ? AND user2_id = ? OR user1_id = ? AND user2_id = ?",friendship.User1ID, friendship.User2ID, friendship.User2ID, friendship.User1ID).
		Delete(models.Friendship{}).Error; err != nil {
		return err
	}
	return nil
}
