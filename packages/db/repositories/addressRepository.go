package repositories

import (
	"github.com/CSVaishakh/QuickHand/packages/db/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AddressRepository struct {
	db *gorm.DB
}

func NewAddressRepository(
	db *gorm.DB,
) *AddressRepository {
	return &AddressRepository{
		db: db,
	}
}

func (repo *AddressRepository) WithTx(tx *gorm.DB) *AddressRepository {
	if tx == nil {
		return repo
	}
	return &AddressRepository{
		db: tx,
	}
}

func (repo *AddressRepository) AddAddress(
	address *models.Address,
) error {
	return repo.db.Create(address).Error
}

func (repo *AddressRepository) UpdateAddress(
	address *models.Address,
) error {
	return repo.db.Save(address).Error
}

func (repo *AddressRepository) GetAddresses(
	UserID uuid.UUID,
) ([]models.Address, error) {
	var addresses []models.Address
	
	err := repo.db.
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
) (models.Address, error) {
	var address models.Address

	err := repo.db.
		Where("address_id = ?", AddressID).
		First(&address).
		Error

	if err != nil {
		return models.Address{}, err
	}

	return address, nil	
}