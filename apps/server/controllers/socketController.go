package controllers

import (
	"log"

	"github.com/CSVaishakh/QuickHand/apps/server/middleware"
	"github.com/CSVaishakh/QuickHand/packages/auth"
	alert "github.com/CSVaishakh/QuickHand/apps/server/services/alertService"
	ss "github.com/CSVaishakh/QuickHand/packages/websockets"

	"github.com/gofiber/contrib/v3/websocket"
	"github.com/gofiber/fiber/v3"
)

type SocketController struct {
	*Controller
	SocketService	*ss.SocketService
	AlertService 	*alert.AlertService
	AuthService  	*auth.AuthService
}

func NewSocketController(
	router 				fiber.Router,
	socketService 		*ss.SocketService,
	alertService 		*alert.AlertService,
	authService 		*auth.AuthService,
) *SocketController {
	return &SocketController{
		Controller:   		NewController(router),
		SocketService: 	socketService,
		AlertService: 		alertService,
		AuthService:  		authService,
	}
}

func (c *SocketController) RegisterRoutes() {
	socketRouter := c.Router.Group("/ws")

	socketRouter.Use(
		middleware.RequireAuthWS(
			c.AuthService,
		),
	)

	socketRouter.Get("/register", websocket.New(c.RegisterSocket))
	socketRouter.Get("/unregister", websocket.New(c.UnregisterSocket))
}

func (c *SocketController) RegisterSocket(conn *websocket.Conn) {
	claims, ok := conn.Locals("claims").(*auth.Claims)
	if !ok {
		_ = conn.WriteJSON(fiber.Map{"error": "unauthorized"})
		_ = conn.Close()
		return
	}

	req := ss.RegisterReq{
		UserID: claims.UserID,
		Conn:   conn,
	}
	
	if err := c.SocketService.Register(req); err != nil {
		log.Println("Failed to register socket:", err.Error())
		_ = conn.Close()
		return
	}

	for {
		if _, _, err := conn.ReadMessage(); err != nil {
			break
		}
	}
}

func (c *SocketController) UnregisterSocket(conn *websocket.Conn) {
	claims, ok := conn.Locals("claims").(*auth.Claims)
	if !ok {
		_ = conn.WriteJSON(fiber.Map{"error": "unauthorized"})
		_ = conn.Close()
		return
	}

	req := ss.UnregisterReq{
		UserID: claims.UserID,
	}

	if err := c.SocketService.Unregister(req); err != nil {
		log.Println("Failed to unregister socket:", err.Error())
		_ = conn.WriteJSON(fiber.Map{"error": err.Error()})
		_ = conn.Close()
		return
	}

	_ = conn.WriteJSON(fiber.Map{"status": "unregistered"})
	_ = conn.Close()
}
