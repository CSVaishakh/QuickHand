package repositories

import (
	"errors"

	"github.com/CSVaishakh/QuickHand/packages/db/models"
	
	"gorm.io/gorm"
)

type ClientRepository struct {
	db *gorm.DB
}

func NewClientRepository(
	db *gorm.DB,
) *ClientRepository {
	return &ClientRepository{
		db: db,
	}
}

func (repo *ClientRepository) WithTx(tx *gorm.DB) *ClientRepository {
	if tx == nil {
		return repo
	}
	return &ClientRepository{
		db: tx,
	}
}

func (repo *ClientRepository) GetByEmail(
	email string,
) (*models.Client, error) {
	var user models.Client

	err := repo.db.
		Table("users").
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
) (*models.Client, error) {
	var user models.Client

	err := repo.db.
		Table("users").
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