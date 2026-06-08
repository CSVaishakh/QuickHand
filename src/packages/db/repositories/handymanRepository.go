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

func (repo *HandymenRepository) CreateUser(
	user *models.Handyman,
	tx *gorm.DB,
) error {
	err := tx.Create(&user.User).Error
	if err != nil {
		return err
	}

	return nil
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

func (repo *HandymenRepository) CheckByEmail(
	email string,
	tx *gorm.DB,
) (bool, error) {
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

func (repo *HandymenRepository) GetUser(
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
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return &user, nil
}
