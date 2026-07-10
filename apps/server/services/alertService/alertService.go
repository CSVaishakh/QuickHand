package alertService

import (
	repo "github.com/CSVaishakh/QuickHand/packages/db/repositories"
	ss "github.com/CSVaishakh/QuickHand/packages/websockets"
)

type AlertService struct {
	alertRepo 			*repo.AlertRepository
	socketService 		*ss.SocketService
}

func NewAlertService (
	alertRepo 			*repo.AlertRepository,
	socketService 		*ss.SocketService,
) *AlertService {
	return &AlertService{
		alertRepo: 			alertRepo,
		socketService: 	socketService,
	}
}

func (s *AlertService) RegisterUserSocket (req RegisterSocketReq) error {
	
	err := s.socketService.Register(req.UserID, req.Conn)
	if err != nil {
		return err
	}
	
	return nil
}

func (s *AlertService) UnregisterUserSocket(req UnregisterSocketReq) error {

	err := s.socketService.Unregister(req.UserID)
	if err != nil {
		return err
	}

	return nil
}

func ( s *AlertService) SendAlert (req SendAlertReq) error {
	err := s.socketService.Send(
		req.UserID, 
		AlertPayload {
			Title:			req.Title, 
			Description: 	req.Description,
		},
	)

	return err
}