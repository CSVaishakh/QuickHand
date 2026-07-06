package main
import (
	"log"
	"os"

	as "github.com/CSVaishakh/QuickHand/src/apps/server/services/addressService"

	ctrs "github.com/CSVaishakh/QuickHand/src/apps/server/controllers"
	
	auth "github.com/CSVaishakh/QuickHand/src/packages/auth/src"
	DB "github.com/CSVaishakh/QuickHand/src/packages/db"
	// ws "github.com/CSVaishakh/QuickHand/src/packages/websockets"
	
	repositories "github.com/CSVaishakh/QuickHand/src/packages/db/repositories"

	"github.com/gofiber/fiber/v3"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(".env.local"); err != nil {
		log.Println("No env file found")
	}

	// Initialize database
	db_url := os.Getenv("DATABASE_URL")
	db := DB.Init(db_url)

	// Repositories
	userRepo 		:= repositories.NewUserRepository(db)
	handymenRepo 	:= repositories.NewHandymenRepository(db)
	clientRepo 		:= repositories.NewClientRepository(db)
	sessionRepo 	:= repositories.NewSessionRepository(db)
	addressRepo 	:= repositories.NewAddressRepository(db)

	// Services
	jwtService := auth.NewJWTService(
		os.Getenv("AUTH_SECRET"),
	)

	authService := auth.NewAuthService(
		userRepo,
		handymenRepo,
		clientRepo,
		sessionRepo,
		jwtService,
		db,
	)

	addressService := as.NewAddressService(
		addressRepo,
	)

	// socketService := ws.NewSocketService()
	// Fiber app
	app := fiber.New()

	// Controllers
	authController := ctrs.NewAuthController(
		app,
		authService,
	)

	addressController := ctrs.NewAddressController(
		app,
		addressService,
		authService,
	)
	
	//Route Regsitrations
	authController.RegisterRoutes()
	addressController.RegisterRoutes()

	// Health check
	app.Get("/", func(c fiber.Ctx) error {
		return c.SendString("Welcome to QuickHand")
	})

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	log.Printf("Server running on :%s", port)

	log.Fatal(
		app.Listen(":" + port),
	)
}