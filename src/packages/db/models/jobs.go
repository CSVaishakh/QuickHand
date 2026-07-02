package models

import (
	"time"

	"github.com/google/uuid"
)

type JobType string

const (
	Plumbing      JobType = "plumbing"
	Electrical    JobType = "electrical"
	Carpentry     JobType = "carpentry"
	Masonry       JobType = "masonry"
	Mechanical    JobType = "mechanical"
	HVAC          JobType = "hvac"
	Landscaping   JobType = "landscaping"
	DeepCleaning  JobType = "deep_cleaning"
)

type HireType string

const (
	DirectHire HireType = "direct_hire"
	BidToGet   HireType = "bid_to_get"
)

type UrgencyLevel string

const (
	Instant            UrgencyLevel = "instant"
	Urgent             UrgencyLevel = "urgent"
	EarliestAvailable  UrgencyLevel = "earliest_available"
	Flexible           UrgencyLevel = "flexible"
)

type Job struct {
	JobID       uuid.UUID      `gorm:"column:job_id;type:uuid;default:gen_random_uuid();primaryKey"`
	ClientID    uuid.UUID     	`gorm:"column:client_id"`
	HandymanID  *uuid.UUID     `gorm:"column:handyman_id"`
	JobType     JobType        `gorm:"column:job_type"`
	HireType    HireType       `gorm:"column:hire_type"`
	Description string         `gorm:"column:description"`
	Budget      float64        `gorm:"column:budget"`
	CreatedAt	time.Time		`gorm:"column:created_at"`
	DeadlineAt  time.Time      `gorm:"column:deadline_at"`
	Urgency     UrgencyLevel   `gorm:"column:urgency"`
}