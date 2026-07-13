package models

import (
	"github.com/google/uuid"
)

type QueuedMessage struct {
	QueueID   uuid.UUID `gorm:"column:queue_id;type:uuid;default:gen_random_uuid();primaryKey"`
	MessageID uuid.UUID `gorm:"column:message_id;type:uuid;unique;not null"`
}

func (QueuedMessage) TableName() string {
	return "message_queue"
}