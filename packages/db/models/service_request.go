package models

import (
    "github.com/google/uuid"
)

type ServiceRequest struct {
    RequestID    uuid.UUID   `gorm:"column:req_id;type:uuid;default:gen_random_uuid();primaryKey"`
    ClientID     uuid.UUID   `gorm:"column:client_id"`
    HandymanID   *uuid.UUID  `gorm:"column:handyman_id"`
    JobID        uuid.UUID   `gorm:"column:job_id"`
    Status       StatusType  `gorm:"column:status"`
}