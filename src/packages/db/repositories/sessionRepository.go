package repositories

import (
	"github.com/CSVaishakh/QuickHand/src/packages/db/models"
	"gorm.io/gorm"
)

type SessionRepository struct {
	db *gorm.DB
}

func NewSessionRepository(
	db *gorm.DB,
) *SessionRepository {
	return &SessionRepository{
		db: db,
	}
}

func (repo *SessionRepository) CreateSession (
	session *models.Session,
	) error {
	return repo.db.Create(session).Error
}