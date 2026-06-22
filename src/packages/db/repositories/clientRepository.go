package repositories

import (
	"errors"

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
	tx *gorm.DB,
) error {

	return tx.Create(&user.User).Error
}

func (repo *ClientRepository) CheckByEmail (
	email string,
	tx *gorm.DB,
) (bool,error){
	var count int64

	res := tx.Raw(
		"SELECT count(*) FROM users WHERE email = ?", 
		email,
	).Scan(&count)
	
	if res.Error != nil {
		return false, res.Error
	}

	return count > 0, nil
}

func (repo *ClientRepository) GetUserByEmail(
	email string,
	tx *gorm.DB,
) (*models.Client, error) {
	var user models.Client

	err := tx.
		Where("email = ?", email).
		Take(&user).
		Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, gorm.ErrRecordNotFound
	}

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (repo *ClientRepository) GetUserByID(
	UserId string,
	tx *gorm.DB,
) (*models.Client, error) {
	var user models.Client

	err := tx.
		Where("user_id = ?", UserId).
		Take(&user).
		Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, gorm.ErrRecordNotFound
	}

	if err != nil {
		return nil, err
	}

	return &user, nil
}