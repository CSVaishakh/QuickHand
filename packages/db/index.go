package db

import (
	"gorm.io/gorm"

	"log"
)

func Init(db_url string) (db *gorm.DB) {

	connection, err := New(db_url)

	if err != nil {
		log.Fatal("database connection failed")
	}

	db = connection

	return db
}
