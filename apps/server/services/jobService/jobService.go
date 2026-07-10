package jobService

import (
	"errors"

	"github.com/CSVaishakh/QuickHand/packages/db/models"
	repo "github.com/CSVaishakh/QuickHand/packages/db/repositories"

	"gorm.io/gorm"
)

type JobService struct {
	jobRepo        *repo.JobRepository
	handymanRepo   *repo.HandymenRepository
	serviceReqRepo *repo.ServiceRequestRepository
	db             *gorm.DB
}

func NewJobService(
	jobRepo *repo.JobRepository,
	handymanRepo *repo.HandymenRepository,
	serviceReqRepo *repo.ServiceRequestRepository,
	db *gorm.DB,
) *JobService {
	return &JobService{
		jobRepo:        jobRepo,
		handymanRepo:   handymanRepo,
		serviceReqRepo: serviceReqRepo,
		db:             db,
	}
}

func (s *JobService) CreateJob(req CreateJobReq) (CreateJobsRes, error) {

	job := models.Job{
		ClientID:    req.ClientID,
		JobType:     req.JobType,
		HireType:    req.HireType,
		Description: req.Description,
		Budget:      req.Budget,
		DeadlineAt:  req.Deadline_At,
		Urgency:     req.Urgency,
	}

	err := s.jobRepo.CreateJob(&job)
	if err != nil {
		return CreateJobsRes{}, err
	}

	return CreateJobsRes{
		Job: job,
	}, nil
}

func (s *JobService) RequestHandymanService(req HandymanServiceReq) error {
	handyman, err := s.handymanRepo.GetByUserID(req.HandymanID)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return ErrInvalidHandyman
	}

	if err != nil {
		return err
	}

	job, err := s.jobRepo.GetJobByJobID(req.JobID)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return ErrInvalidJob
	}

	if err != nil {
		return err
	}

	if job.JobType != handyman.Type.MapJobType() {
		return ErrTypeMismatch
	}

	if job.ClientID != req.ClientID {
		return ErrUnauthorized
	}

	err = s.db.Transaction(func(tx *gorm.DB) error {
		jobRepoTx := s.jobRepo.WithTx(tx)
		reqRepoTx := s.serviceReqRepo.WithTx(tx)

		err = jobRepoTx.MarkAsRequested(req.JobID)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ErrInvalidJob
		}
		if err != nil {
			return err
		}

		serviceReq := models.ServiceRequest{
			ClientID:   req.ClientID,
			HandymanID: &req.HandymanID,
			JobID:      req.JobID,
			Status:     models.Requested,
		}

		err = reqRepoTx.CreateServiceRequest(&serviceReq)
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return err
	}

	return nil
}
