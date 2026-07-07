package db

import "gorm.io/gorm"

func Transaction (
	db *gorm.DB, 
	fn func(tx *gorm.DB) error,
) error {
	return	db.Transaction(fn) 
}
