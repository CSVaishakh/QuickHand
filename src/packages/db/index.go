package db

import (
	"github.com/CSVaishakh/QuickHand/src/packages/db/src"
	"gorm.io/gorm"

	"log"
)

func Init(db_url string) (db *gorm.DB) {

	connection, err := src.New(db_url)

	if err != nil {
		log.Fatal("database connection failed")
	}

	db = connection

	return db
}
