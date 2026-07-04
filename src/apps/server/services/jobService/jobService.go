package jobService

import (
	"github.com/CSVaishakh/QuickHand/src/packages/db/models"
	repo "github.com/CSVaishakh/QuickHand/src/packages/db/repositories"

	"gorm.io/gorm"
)

type JobService struct{
	jobRepo repo.JobRepository
	db *gorm.DB
}

func NewJobService (
	jobRepo *repo.JobRepository,
	db *gorm.DB,
) *JobService {
	return &JobService{
		jobRepo: *jobRepo,
		db: db,
	}
}

func(s *JobService) CreateJob (req CreateJobReq) (CreateJobsRes, error) {
	
	job:= models.Job{
		ClientID: req.ClientID,
		JobType: req.JobType,
		HireType: req.HireType,
		Description: req.Description,
		Budget: req.Budget,
		DeadlineAt: req.Deadline_At,
		Urgency: req.Urgency,
	}

	err := s.jobRepo.CreateJob(&job, s.db)
	if err != nil {
		return CreateJobsRes{}, err
	}

	return CreateJobsRes{
		Job: job,
	}, nil
}