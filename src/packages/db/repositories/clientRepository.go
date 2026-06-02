package repositories

import (
	"github.com/CSVaishakh/QuickHand/src/packages/db/models"
	"gorm.io/gorm"
)

type ClientRepository struct {
	db *gorm.DB
}

func NewClientRepository (
	db *gorm.DB,
) *ClientRepository {
	return  &ClientRepository{
		db: db,
	}
}

func (repo *ClientRepository) CreateUser (
		user *models.Client,
	) error {

	return repo.db.Create(&user.User).Error
}