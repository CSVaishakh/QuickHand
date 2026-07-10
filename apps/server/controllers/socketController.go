package controllers

import (
	"log"

	"github.com/CSVaishakh/QuickHand/apps/server/middleware"
	alert "github.com/CSVaishakh/QuickHand/apps/server/services/alertService"
	"github.com/CSVaishakh/QuickHand/packages/auth"

	"github.com/gofiber/contrib/v3/websocket"
	"github.com/gofiber/fiber/v3"
)

type SocketController struct {
	*Controller
	AlertService *alert.AlertService
	AuthService  *auth.AuthService
}

func NewSocketController(
	router fiber.Router,
	alertService *alert.AlertService,
	authService *auth.AuthService,
) *SocketController {
	return &SocketController{
		Controller:   NewController(router),
		AlertService: alertService,
		AuthService:  authService,
	}
}

func (c *SocketController) RegisterRoutes() {
	socketRouter := c.Router.Group("/ws")

	socketRouter.Use(
		middleware.RequireAuthWS(
			c.AuthService,
		),
	)

	socketRouter.Get("/register-alert-socket", websocket.New(c.HandleAlertSocket))
}

func (c *SocketController) HandleAlertSocket(conn *websocket.Conn) {
	claims, ok := conn.Locals("claims").(*auth.Claims)
	if !ok {
		_ = conn.WriteJSON(fiber.Map{"error": "unauthorized"})
		_ = conn.Close()
		return
	}

	req := alert.RegisterSocketReq{
		UserID: claims.UserID,
		Conn:   conn,
	}
	if err := c.AlertService.RegisterUserSocket(req); err != nil {
		log.Println("Failed to register socket:", err.Error())
		_ = conn.Close()
		return
	}

	defer func() {
		_ = c.AlertService.UnregisterUserSocket(alert.UnregisterSocketReq{
			UserID: claims.UserID,
			Conn:   conn,
		})
	}()

	for {
		if _, _, err := conn.ReadMessage(); err != nil {
			break
		}
	}
}
