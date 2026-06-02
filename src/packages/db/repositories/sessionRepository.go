package repositories

import (
	"github.com/CSVaishakh/QuickHand/src/packages/db/models"
	"gorm.io/gorm"
)

type sessionRepository struct {
	db *gorm.DB
}

func NewSessionRepository(
	db *gorm.DB,
) *sessionRepository {
	return &sessionRepository{
		db: db,
	}
}

func (repo *sessionRepository) CreateSession (
	session *models.Session,
	) error {
	return repo.db.Create(session).Error
}