package repositories

import (
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

func (repo *HandymenRepository) CreateUser (
	user *models.Handyman,
) error {

	return repo.db.Transaction(
		func(tx *gorm.DB) error {

			

			err:= tx.Create(&user.User).Error
			if err != nil {
				return err
			}

			err = tx.Exec(
					"INSERT INTO handymen (user_id, type) VALUES (?, ?)",
					user.UserID,
					user.Type,
				).Error

			if err!= nil {
				return err
			}

			return nil
		},
	)
		
}