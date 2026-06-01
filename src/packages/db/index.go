package db

import (
	"github.com/CSVaishakh/QuickHand/src/packages/db/src"
	"gorm.io/gorm"

	"log"
	"os"
)

var db *gorm.DB;

func Init () {

	db_url := os.Getenv("DATABASE_URL")
	connection, err :=  src.New(db_url)

	if err != nil {
		log.Fatal("databse connection failed")
	}

	db = connection;
}