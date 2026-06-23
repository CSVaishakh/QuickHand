package models

import (
	"time"

	"github.com/google/uuid"
)

type Session struct {
	ID 			uuid.UUID `gorm:"column:id;type:uuid;default:gen_random_uuid();primaryKey"`

	UserID 		uuid.UUID `gorm:"column:user_id"`
	TokenHash 	string `gorm:"column:token_hash"`
	Revoked 	bool `gorm:"column:revoked"`

	CreatedAt 	time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt 	time.Time `gorm:"column:updated_at;autoUpdateTime"`
}