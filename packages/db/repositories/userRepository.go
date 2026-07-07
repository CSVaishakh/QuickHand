package repositories

import (
	"errors"
	
	"github.com/CSVaishakh/QuickHand/packages/db/models"
	
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(
	db *gorm.DB,
) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (repo *UserRepository) WithTx(tx *gorm.DB) *UserRepository {
	if tx == nil {
		return repo
	}
	return &UserRepository{
		db: tx,
	}
}

func (repo *UserRepository) CreateUser(
	user *models.User,
) error {
	err := repo.db.Create(&user).Error
	if err != nil {
		return err
	}

	return nil
}

func (repo *UserRepository) CheckByEmail(
	email string,
) (bool, error) {
	var count int64

	res := repo.db.Raw(
		"SELECT count(*) FROM users WHERE email = ?",
		email,
	).Scan(&count)

	if res.Error != nil {
		return false, res.Error
	}

	return count > 0, nil
}

func (repo *UserRepository) CheckByUserID(
	UserId string,
) (bool, error) {
	var count int64

	res := repo.db.Raw(
		"SELECT count(*) FROM users WHERE user_id = ?",
		UserId,
	).Scan(&count)

	if res.Error != nil {
		return false, res.Error
	}

	return count > 0, nil
}

func (repo *UserRepository) ResetPassword(
	HashedPass []byte,
	Email string,
) error {
	var user models.User
	err := repo.db.Raw(
		"UPDATE users SET password_hash = ? where email = ?",
		HashedPass, Email,
	).Scan(&user)

	if errors.Is(err.Error, gorm.ErrRecordNotFound) {
		return gorm.ErrRecordNotFound
	}
	
	if err.Error != nil {
		return err.Error
	}

	return nil
}