package repositories

import (
	"github.com/CSVaishakh/QuickHand/src/packages/db/models"
	"gorm.io/gorm"
)

type clientRepository struct {
	userRepository
}

func (repo *clientRepository) CreateUser (
		db *gorm.DB,
		user *models.Client,
	) error {

	return repo.userRepository.CreateUser(db, &user.User)
}