package repositories

import (
	"github.com/CSVaishakh/QuickHand/src/packages/db/models"
	"gorm.io/gorm"
)

type sessionRepository struct {}

func (repo *sessionRepository) CreateSession (
	db *gorm.DB,
	session *models.Session,
	) error {
	return db.Create(session).Error
}