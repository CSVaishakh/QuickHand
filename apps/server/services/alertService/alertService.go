package alertService

import (
	repo "github.com/CSVaishakh/QuickHand/packages/db/repositories"
	ss "github.com/CSVaishakh/QuickHand/packages/websockets"
	"gorm.io/gorm"
)

import models "github.com/CSVaishakh/QuickHand/packages/db/models"

type AlertService struct {
	alertRepo 		*repo.AlertRepository
	socketService 	*ss.SocketService
	db 				*gorm.DB
}

func NewAlertService (
	alertRepo 		*repo.AlertRepository,
	socketService 	*ss.SocketService,
	db 				*gorm.DB,
) *AlertService {
	return &AlertService{
		alertRepo: 			alertRepo,
		socketService: 	socketService,
		db: 					db,
	}
}

func (s *AlertService) SendAlert(req SendAlertReq) error {
	err := s.db.Transaction(func(tx *gorm.DB) error {
		if err := s.socketService.Send(
			req.UserID,
			AlertPayload{
				Title:       req.Title,
				Description:  req.Description,
			},
		); err != nil {
			return err
		}

		alert := &models.Alert{
			UserID:   req.UserID,
			Title:     req.Title,
			Message:   req.Description,
			IsRead:    false,
		}
		
		if err := s.alertRepo.WithTx(tx).SaveAlert(alert); err != nil {
			return ErrSavingAlert
		}
		
		return nil
	})
		
	return err
}