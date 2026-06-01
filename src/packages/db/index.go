package db

import (

	"github.com/CSVaishakh/QuickHand/src/packages/db/src"

	"log"
	"os"
)

var DB *src.DB;

func Init () {

	db_url := os.Getenv("DATABASE_URL")
	connection, err :=  src.New(db_url)

	if err != nil {
		log.Fatal("databse connection failed")
	}

	DB = connection;
}