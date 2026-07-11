package main
import (
	"log"
	"os"

	as "github.com/CSVaishakh/QuickHand/apps/server/services/addressService"
	alert "github.com/CSVaishakh/QuickHand/apps/server/services/alertService"
	job "github.com/CSVaishakh/QuickHand/apps/server/services/jobService"

	ctrs "github.com/CSVaishakh/QuickHand/apps/server/controllers"
	
	auth "github.com/CSVaishakh/QuickHand/packages/auth"
	DB "github.com/CSVaishakh/QuickHand/packages/db"
	ws "github.com/CSVaishakh/QuickHand/packages/websockets"
	
	repositories "github.com/CSVaishakh/QuickHand/packages/db/repositories"

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
	jobRepo         := repositories.NewJobRepository(db)
	serviceReqRepo := repositories.NewServiceRequestRepository(db)
	alertRepo       := repositories.NewAlertRepository(db)

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

	socketService := ws.NewSocketService()

	alertService := alert.NewAlertService(
		alertRepo,
		socketService,
	)

	jobService := job.NewJobService(
		jobRepo,
		handymenRepo,
		serviceReqRepo,
		db,
	)

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

	socketController := ctrs.NewSocketController(
		app,
		socketService,
		alertService,
		authService,
	)

	jobController := ctrs.NewJobController(
		app,
		jobService,
		authService,
		alertService,
	)
	
	//Route Registrations
	authController.RegisterRoutes()
	addressController.RegisterRoutes()
	socketController.RegisterRoutes()
	jobController.RegisterRoutes()

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