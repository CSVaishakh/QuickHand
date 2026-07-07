package models

import (
	"time"

	"github.com/google/uuid"
)

type UserRole string

const (
	ClientRole   	UserRole = "client"
	HandymanRole 	UserRole = "handyman"
)

type User struct {
	UserID           	uuid.UUID `gorm:"column:user_id;type:uuid;default:gen_random_uuid();primaryKey"`
	FirstName    		string    `gorm:"column:first_name;size:25;not null"`
	LastName     		string    `gorm:"column:last_name;size:25;not null"`
	Email        		string    `gorm:"column:email;size:255;unique;not null"`
	PasswordHash 		string    `gorm:"column:password_hash;type:text;not null"`
	Role         		UserRole  `gorm:"column:role;type:user_role;not null"`
	PhoneNumber  		string    `gorm:"column:phone_number;type:text;not null"`
	Img          		*string   `gorm:"column:img;type:text"`
	CreatedAt    		time.Time `gorm:"column:created_at;autoCreateTime"`
}