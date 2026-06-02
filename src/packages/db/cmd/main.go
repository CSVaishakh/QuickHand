package main

import (
	"log"

	dbmigrate "github.com/CSVaishakh/QuickHand/src/packages/db/migrate"
)

func main() {
	if err := dbmigrate.Up(); err != nil {
		log.Fatal(err)
	}
}