package repositories

import (
	"errors"

	"github.com/CSVaishakh/QuickHand/src/packages/db/models"
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

func (repo *HandymenRepository) AddHandymenType(
	user *models.Handyman,
	tx *gorm.DB,
) error {
	err := tx.Exec(
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
	tx *gorm.DB,
) (*models.Handyman, error) {
	var user models.Handyman

	err := tx.
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
	tx *gorm.DB,
) (*models.Handyman, error) {
	var user models.Handyman

	err := tx.
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
