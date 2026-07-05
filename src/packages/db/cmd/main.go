package main

import (
	"fmt"
	"log"
	"os"

	dbmigrate "github.com/CSVaishakh/QuickHand/src/packages/db/migrate"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		if err := dbmigrate.Up(); err != nil {
			log.Fatal(err)
		}
		return
	}

	switch args[0] {
	case "up":
		if err := dbmigrate.Up(); err != nil {
			log.Fatal(err)
		}
	case "down":
		if err := dbmigrate.Down(); err != nil {
			log.Fatal(err)
		}
	case "version":
		if err := dbmigrate.Version(); err != nil {
			log.Fatal(err)
		}
	case "force":
		if len(args) < 2 {
			log.Fatal("usage: migrate force <version>")
		}
		var version int
		if _, err := fmt.Sscanf(args[1], "%d", &version); err != nil {
			log.Fatal(err)
		}
		if err := dbmigrate.Force(version); err != nil {
			log.Fatal(err)
		}
	case "create":
		if len(args) < 2 {
			log.Fatal("usage: migrate create <name>")
		}
		if err := dbmigrate.Create(args[1]); err != nil {
			log.Fatal(err)
		}
	default:
		log.Fatalf("unknown command: %s", args[0])
	}
}