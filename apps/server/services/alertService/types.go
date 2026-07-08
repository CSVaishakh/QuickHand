package alertService
import (
	"github.com/gofiber/contrib/v3/websocket"
	"github.com/google/uuid"
)

type RegisterConnectionReq struct {
	UserID 	uuid.UUID
	Conn 		*websocket.Conn
}

type UnregisterConnectionReq struct {
	UserID 	uuid.UUID
	Conn 		*websocket.Conn
}

type SendAlertReq struct {
	UserID 			uuid.UUID
	Title				string
	Description 	string
}

type AlertPayload struct {
    Title       string `json:"title"`
    Description string `json:"description"`
}