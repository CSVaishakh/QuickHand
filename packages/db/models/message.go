package models

import (
	"time"

	"github.com/google/uuid"
)

type MessageStatus string

const (
	MessageStatusQueued    MessageStatus = "queued"
	MessageStatusDelivered MessageStatus = "delivered"
)

type Message struct {
	MessageID uuid.UUID     `gorm:"column:message_id;type:uuid;default:gen_random_uuid();primaryKey"`
	SenderID  uuid.UUID     `gorm:"column:sender_id;type:uuid;not null"`
	ReciverID uuid.UUID     `gorm:"column:reciver_id;type:uuid;not null"`
	Message   string        `gorm:"column:message;type:text;not null"`
	Status    MessageStatus `gorm:"column:status;type:status_types;default:delivered;not null"`
	SentAt    time.Time     `gorm:"column:sent_at;type:timestamptz;default:now();not null"`
}