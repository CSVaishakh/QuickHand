package repositories

import (

	"github.com/CSVaishakh/QuickHand/src/packages/db/models"

	"github.com/google/uuid"

	"gorm.io/gorm"
	
)

type AddressRepository struct {
	db *gorm.DB
}

func NewAddressRepository (
	db *gorm.DB,
) *AddressRepository {
	return &AddressRepository{
		db: db,
	}
}

func (repo *AddressRepository) AddAddress (
	address *models.Address,
	tx *gorm.DB,
) error {
	return tx.Create(address).Error
}

func (repo *AddressRepository) UpdateAddress(
	address *models.Address,
	tx *gorm.DB,
) error {
	return tx.Save(address).Error
}

func (repo *AddressRepository) GetAddresses(
	userID string,
	tx *gorm.DB,
) ([]models.Address, error) {
	var addresses []models.Address

	id, err := uuid.Parse(userID)
	if err != nil {
		return nil, err
	}

	err = tx.
		Where("user_id = ?", id).
		Find(&addresses).
		Error

	if err != nil {
		return nil, err
	}

	return addresses, nil
}