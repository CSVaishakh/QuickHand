package repositories

import (
	"errors"

	"github.com/CSVaishakh/QuickHand/packages/db/models"
	
	"gorm.io/gorm"
)

type HandymenRepository struct {
	db *gorm.DB
}

func NewHandymenRepository(
	db *gorm.DB,
) *HandymenRepository {
	return &HandymenRepository{
		db: db,
	}
}

func (repo *HandymenRepository) WithTx(tx *gorm.DB) *HandymenRepository {
	if tx == nil {
		return repo
	}
	return &HandymenRepository{
		db: tx,
	}
}

func (repo *HandymenRepository) AddHandymenType(
	user *models.Handyman,
) error {
	err := repo.db.Exec(
		"INSERT INTO handymen (user_id, type) VALUES (?, ?)",
		user.UserID,
		user.Type,
	).Error

	if err != nil {
		return err
	}
	return nil
}

func (repo *HandymenRepository) GetByEmail(
	email string,
) (*models.Handyman, error) {
	var user models.Handyman

	err := repo.db.
		Table("users").
		Select(`
			users.*,
			handymen.type AS type
		`).
		Joins(`
			JOIN handymen
			ON handymen.user_id = users.user_id
		`).
		Where("users.email = ?", email).
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

func (repo *HandymenRepository) GetByUserID(
	UserId string,
) (*models.Handyman, error) {
	var user models.Handyman

	err := repo.db.
		Table("users").
		Select(`
			users.*,
			handymen.type AS type
		`).
		Joins(`
			JOIN handymen
			ON handymen.user_id = users.user_id
		`).
		Where("users.user_id = ?", UserId).
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

func (repo *HandymenRepository) GetAllByType(
	handymanType models.HandymanType,
) ([]models.Handyman, error) {

	var handymen []models.Handyman

	err := repo.db.
		Table("handymen").
		Select(`
			users.*,
			handymen.type
		`).
		Joins("JOIN users ON users.user_id = handymen.user_id").
		Where("handymen.type = ?", handymanType).
		Find(&handymen).
		Error

	return handymen, err
}