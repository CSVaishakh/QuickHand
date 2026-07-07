package models

import (
   "time"
   "github.com/google/uuid"
)

type Alert struct {
   UserID    	uuid.UUID `gorm:"primaryKey;type:uuid"`
   CreatedAt 	time.Time `gorm:"primaryKey"`
   Title     	string    `gorm:"not null"`
   Message   	string    `gorm:"not null"`
   IsRead    	bool      `gorm:"default:false;not null"`
}