package websockets

import (
	"github.com/gofiber/contrib/v3/websocket"
	"github.com/google/uuid"
)

type RegisterReq struct {
	UserID 	uuid.UUID
	Conn 		*websocket.Conn
}

type UnregisterReq struct {
	UserID	uuid.UUID
}