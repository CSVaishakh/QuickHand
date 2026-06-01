package repositories

import (
	"github.com/CSVaishakh/QuickHand/src/packages/db/models"
	"gorm.io/gorm"
)

type handymenRepository struct {
	userRepository
}

func (repo *handymenRepository) CreateUser (
	db *gorm.DB,
	user *models.Handyman,
	) error {

	return db.Transaction(
		func(tx *gorm.DB) error {
	
			err:= repo.userRepository.CreateUser(tx, &user.User)
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