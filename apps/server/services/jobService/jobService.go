package jobService

import (
	"github.com/CSVaishakh/QuickHand/packages/db/models"
	repo "github.com/CSVaishakh/QuickHand/packages/db/repositories"
)

type JobService struct{
	jobRepo repo.JobRepository
}

func NewJobService (
	jobRepo *repo.JobRepository,
) *JobService {
	return &JobService{
		jobRepo: *jobRepo,
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

	err := s.jobRepo.CreateJob(&job)
	if err != nil {
		return CreateJobsRes{}, err
	}

	return CreateJobsRes{
		Job: job,
	}, nil
}