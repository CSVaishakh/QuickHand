package repositories

import (
	"github.com/CSVaishakh/QuickHand/packages/db/models"

	"gorm.io/gorm"
)

type ServiceRequestRepository struct {
	db *gorm.DB
}

func NewServiceRequestRepository(
	db *gorm.DB,
) *ServiceRequestRepository {
	return &ServiceRequestRepository{
		db: db,
	}
}

func (repo *ServiceRequestRepository) WithTx(tx *gorm.DB) *ServiceRequestRepository {
	if tx == nil {
		return repo
	}
	return &ServiceRequestRepository{
		db: tx,
	}
}

func (repo *ServiceRequestRepository) CreateServiceRequest(
	serviceRequest *models.ServiceRequest,
) error {
	return repo.db.Create(serviceRequest).Error
}