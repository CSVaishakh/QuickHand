package repositories

import (
	"github.com/CSVaishakh/QuickHand/src/packages/db/models"
	
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type JobRepository struct {
	db *gorm.DB
}

func NewJobRepository (
	db *gorm.DB,
) *JobRepository {
	return &JobRepository{
		db: db,
	}
}

func (repo *JobRepository) CreateJob (
	job 	*models.Job,
	tx *gorm.DB,
) error {
	return tx.Create(job).Error
}

func (reop *JobRepository) HandymanRejected(
	jobId uuid.UUID,
	tx *gorm.DB,
) error {
	err := tx.Raw(`
		UPDATE jobs
		SET status = "rejected"
		WHERE handyman_id is NULL
		AND job_id = ?
	`, jobId)

	if err != nil {
		return err.Error
	}

	if tx.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}	

	return nil
}

func (repo *JobRepository) AssignHandyman (
	handymanId uuid.UUID,
	jobID uuid.UUID,
	jobType models.JobType,
	tx *gorm.DB,
)(models.Job, error){
	var job models.Job

	err := tx.Raw(`
		UPDATE jobs
		SET handyman_id = ?, status = "hired"
		WHERE handyman_id is NULL
		AND job_id = ?
		AND job_type = ?
		RETURNING *
	`, handymanId, jobID, jobType).Scan(&job).Error
	
	if err != nil {
		return models.Job{}, err
	}

	if tx.RowsAffected == 0 {
		return models.Job{}, gorm.ErrRecordNotFound
	}

	return job, nil
}

func (repo *JobRepository) GetUserJobs (
	UserID uuid.UUID,
	tx *gorm.DB,
)([]models.Job, error){
	var jobs []models.Job

	err := tx.Raw(`
		SELECT * 
		FROM jobs
		WHERE client_id = ?
			OR handyman_id = ?
	`, UserID, UserID).Scan(&jobs).Error

	if err != nil {
		return []models.Job{},err
	}

	return jobs, nil
}

func (repo *JobRepository) GetJobsToQuote (
	hiringType models.HireType,
	jobType models.JobType,
	tx *gorm.DB,
)([]models.Job, error){
	var jobsToQuote []models.Job

	err := tx.Raw(`
		SELECT * 
		FROM jobs
		WHERE hire_type = ?
		AND job_type = ?
	`, hiringType, jobType).Scan(&jobsToQuote).Error

	if err != nil {
		return []models.Job{}, err
	}

	return jobsToQuote, nil
}