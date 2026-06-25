package repositories

import (
	"errors"

	"github.com/CSVaishakh/QuickHand/src/packages/db/models"
	"github.com/google/uuid"
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
	tx *gorm.DB,
	) error {
	return tx.Create(session).Error
}

func (repo *SessionRepository) GetSession (
	tokenHash string,
) (*models.Session, error) {
	var session models.Session

	err := repo.db.
			Where("token_hash = ?", tokenHash).
			First(&session).
			Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return &session, nil
}

func (repo *SessionRepository) RevokeSession (
	tokenHash string,
) (*models.Session, error) {
	var session models.Session

	err := repo.db.Raw("UPDATE sessions SET revoked = TRUE WHERE token_hash = ? RETURNING *",
		tokenHash,
	).Scan(&session).Error

	if err != nil {
		return nil, err
	}

	if session.SessionID == uuid.Nil {
    	return nil, errors.New("session not found")
	}

	return &session, nil
}