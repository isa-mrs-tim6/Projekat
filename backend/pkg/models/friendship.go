package models

import "github.com/jinzhu/gorm"

type Friendship struct {
	gorm.Model
	Status  string // Requested, Accepted, Denied
	User1   *User  `gorm:"foreignkey:User1ID"`
	User1ID uint
	User2   *User `gorm:"foreignkey:User2ID"`
	User2ID uint
}
