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
	UserID uuid.UUID,
	tx *gorm.DB,
) ([]models.Address, error) {
	var addresses []models.Address
	
	err := tx.
		Where("user_id = ?", UserID).
		Find(&addresses).
		Error

	if err != nil {
		return nil, err
	}

	return addresses, nil
}

func (repo *AddressRepository) GetByAddressID(
	AddressID uuid.UUID,
	tx *gorm.DB,
)(models.Address, error) {
	var address models.Address

	err := tx.
		Where("address_id = ?", AddressID).
		First(&address).
		Error

	if err != nil {
		return models.Address{}, err
	}

	return address, nil	
}