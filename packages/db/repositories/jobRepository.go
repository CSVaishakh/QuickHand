package repositories

import (
	"github.com/CSVaishakh/QuickHand/packages/db/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type JobRepository struct {
	db *gorm.DB
}

func NewJobRepository(
	db *gorm.DB,
) *JobRepository {
	return &JobRepository{
		db: db,
	}
}

func (repo *JobRepository) WithTx(tx *gorm.DB) *JobRepository {
	if tx == nil {
		return repo
	}
	return &JobRepository{
		db: tx,
	}
}

func (repo *JobRepository) CreateJob(
	job *models.Job,
) error {
	return repo.db.Create(job).Error
}

func (repo *JobRepository) AssignHandyman(
	handymanId uuid.UUID,
	jobID uuid.UUID,
	jobType models.JobType,
) (models.Job, error) {
	var job models.Job

	res := repo.db.Raw(`
		UPDATE jobs
		SET handyman_id = ?, status = 'hired'
		WHERE handyman_id IS NULL
		AND job_id = ?
		AND job_type = ?
		RETURNING *
	`, handymanId, jobID, jobType).Scan(&job)

	if res.Error != nil {
		return models.Job{}, res.Error
	}

	if res.RowsAffected == 0 {
		return models.Job{}, gorm.ErrRecordNotFound
	}

	return job, nil
}

func (repo *JobRepository) HandymanRejected(
	jobID uuid.UUID,
) error {
	res := repo.db.Exec(`
		UPDATE jobs
		SET status = 'rejected'
		WHERE handyman_id IS NULL
		AND job_id = ?
	`, jobID)

	if res.Error != nil {
		return res.Error
	}

	if res.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

func (repo *JobRepository) GetUserJobs(
	userID uuid.UUID,
) ([]models.Job, error) {
	var jobs []models.Job

	err := repo.db.Raw(`
		SELECT *
		FROM jobs
		WHERE client_id = ?
			OR handyman_id = ?
	`, userID, userID).Scan(&jobs).Error

	if err != nil {
		return []models.Job{}, err
	}

	return jobs, nil
}

func (repo *JobRepository) GetJobsToQuote(
	hiringType models.HireType,
	jobType models.JobType,
) ([]models.Job, error) {
	var jobsToQuote []models.Job

	err := repo.db.Raw(`
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

func (repo *JobRepository) GetJobByJobID(
	jobID uuid.UUID,
) (*models.Job, error) {
	var job models.Job

	res := repo.db.Raw(`
		SELECT *
		FROM jobs
		WHERE job_id = ?
	`, jobID).Scan(&job)

	if res.Error != nil {
		return &models.Job{}, res.Error
	}

	if res.RowsAffected == 0 {
		return &models.Job{}, gorm.ErrRecordNotFound
	}

	return &job, nil
}

func (repo *JobRepository) MarkAsRequested(
	jobID uuid.UUID,
) error {
	res := repo.db.Exec(`
		UPDATE jobs
		SET status = 'requested'
		WHERE handyman_id IS NULL
		AND job_id = ?
	`, jobID)

	if res.Error != nil {
		return res.Error
	}

	if res.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}
