package models

import "github.com/google/uuid"

type Address struct {
	AddressID 	uuid.UUID `gorm:"column:address_id;type:uuid;default:gen_random_uuid();primaryKey"`

	UserID 		uuid.UUID `gorm:"column:user_id;type:uuid;not null"`

	HouseNo 	string `gorm:"column:house_no;size:50;not null"`
	Street  	string `gorm:"column:street;size:255;not null"`
	City    	string `gorm:"column:city;size:100;not null"`
	State  		string `gorm:"column:state;size:100;not null"`
	Country 	string `gorm:"column:country;size:100;not null"`
	Pincode 	string `gorm:"column:pincode;size:20;not null"`
}