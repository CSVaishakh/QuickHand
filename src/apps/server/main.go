package main

import (
	"log"
	"os"

	"github.com/CSVaishakh/QuickHand/src/packages/db"

	"github.com/gofiber/fiber/v3"
	"github.com/joho/godotenv"
)

func main() {
	//loading enviornment variables
	err:= godotenv.Load(".env.local");
	if err != nil {
		log.Println("No env fil found")
	}

	//starting the databse service
	db.Init()

	//creating the fiber app instance
	app:= fiber.New();
	
	//loading the post from env
	port:= os.Getenv("PORT")

	app.Get("/", func(c fiber.Ctx) error {
		//logging server is working
		return c.Res().SendString("Welcome to QuickHand");
	});

	log.Fatal(app.Listen(":" + port))
} 