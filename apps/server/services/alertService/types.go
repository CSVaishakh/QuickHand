package alertService
import (
	"github.com/gofiber/contrib/v3/websocket"
	"github.com/google/uuid"
)

type RegisterSocketReq struct {
	UserID 	uuid.UUID
	Conn 		*websocket.Conn
}

type UnregisterSocketReq struct {
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