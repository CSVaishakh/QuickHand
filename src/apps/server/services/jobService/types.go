package jobService

import (
	"time"

	"github.com/CSVaishakh/QuickHand/src/packages/db/models"
	"github.com/google/uuid"
)

type CreateJobReq struct {
	ClientID 	uuid.UUID
	JobType 		models.JobType
	HireType 	models.HireType
	Description	string
	Budget		float64
	Deadline_At time.Time
	Urgency 		models.UrgencyLevel
}

type CreateJobsRes struct {
	Job models.Job
}