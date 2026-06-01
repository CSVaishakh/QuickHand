package repositories

import (
	"github.com/CSVaishakh/QuickHand/src/packages/db/models"
	"gorm.io/gorm"
)

type userRepository struct {}

func (repo *userRepository) CreateUser (
	db *gorm.DB,
	user *models.User,
	) error {
	return db.Create(user).Error
}