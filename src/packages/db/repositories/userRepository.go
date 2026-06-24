package repositories

import (
	"github.com/CSVaishakh/QuickHand/src/packages/db/models"
	"gorm.io/gorm"
)


type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository (
	db *gorm.DB,
) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (repo *UserRepository) CreateUser(
	user *models.User,
	tx *gorm.DB,
) error {
	err := tx.Create(&user).Error
	if err != nil {
		return err
	}

	return nil
}

func (repo *UserRepository) CheckByEmail(
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