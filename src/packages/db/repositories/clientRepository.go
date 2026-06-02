package repositories

import (
	"github.com/CSVaishakh/QuickHand/src/packages/db/models"
	"gorm.io/gorm"
)

type clientRepository struct {
	db *gorm.DB
}

func NewClientRepository (
	db *gorm.DB,
) *clientRepository {
	return  &clientRepository{
		db: db,
	}
}

func (repo *clientRepository) CreateUser (
		user *models.Client,
	) error {

	return repo.db.Create(&user.User).Error
}