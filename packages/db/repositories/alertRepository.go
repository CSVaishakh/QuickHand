package repositories

import (
	"github.com/CSVaishakh/QuickHand/packages/db/models"
	"gorm.io/gorm"
)

type AlertRepository struct {
	db *gorm.DB
}

func NewAlertRepository(
	db *gorm.DB,
) *AlertRepository {
	return &AlertRepository{
		db: db,
	}
}

func (repo *AlertRepository) WithTx(tx *gorm.DB) *AlertRepository {
	if tx == nil {
		return repo
	}
	return &AlertRepository{
		db: tx,
	}
}

func (repo *AlertRepository) SaveAlert(
	alert *models.Alert,
) error {
	return repo.db.Create(alert).Error
}