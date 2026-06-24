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

func (repo *ClientRepository) GetByEmail(
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

func (repo *ClientRepository) GetByUserID(
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