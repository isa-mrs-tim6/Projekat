package postgre

import "github.com/jinzhu/gorm"

type Store struct {
	*gorm.DB
}
