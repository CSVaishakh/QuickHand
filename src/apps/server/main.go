package main

import (
	"log"
	"os"

	DB "github.com/CSVaishakh/QuickHand/src/packages/db"
	repositories "github.com/CSVaishakh/QuickHand/src/packages/db/repositories"
	auth "github.com/CSVaishakh/QuickHand/src/packages/auth/src"


	"github.com/gofiber/fiber/v3"
	"github.com/joho/godotenv"
)

func main() {
	//loading environment variables
	err := godotenv.Load(".env.local")
	if err != nil {
		log.Println("No env fil found")
	}

	//establishing the db connection
	db := DB.Init()
	
	//creating the repository instances
	handymenRepo := repositories.NewHandymenRepository(db)
	clientRepo := repositories.NewClientRepository(db)
	sessionRepo := repositories.NewSessionRepository(db)

	// creating the auth serivce instance
	jwtSerivce := auth.NewJWTService(os.Getenv("AUTH_SECRET"))
	authService := auth.NewAuthService(handymenRepo, clientRepo, sessionRepo, jwtSerivce,db)

	//creating the fiber app instance
	app := fiber.New()

	//loading the post from env
	port := os.Getenv("PORT")

	app.Get("/", func(c fiber.Ctx) error {
		//logging server is working
		return c.Res().SendString("Welcome to QuickHand")
	})

	log.Fatal(app.Listen(":" + port))
}
