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

func (repo *ClientRepository) GetByEmail (
	email string,
) (bool,error){
	var count int64

	res := repo.db.Raw(
		"SELECT count(*) FROM users WHERE email = ?", 
		email,
	).Scan(&count)
	
	if res.Error != nil {
		return false, res.Error
	}

	return count > 0, nil
}