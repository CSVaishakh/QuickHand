package websockets

import (
	"sync"

	"github.com/gofiber/contrib/v3/websocket"
	"github.com/google/uuid"
)

type SocketService struct {
	sockets map[uuid.UUID]*websocket.Conn
	mu sync.RWMutex
}

func NewSocketService() *SocketService {
	return &SocketService{
		sockets: make(map[uuid.UUID]*websocket.Conn),
	}
}

func(ss *SocketService) Register (req RegisterReq) error {

	if req.UserID == uuid.Nil {
		return ErrEmptyUserId
	}

	if req.Conn == nil {
		return ErrEmptyConnectionReference
	}
	
	ss.mu.Lock()
	defer ss.mu.Unlock()

	if oldConn, exists := ss.sockets[req.UserID]; exists {
		_ = oldConn.Close()
	} 

	ss.sockets[req.UserID] = req.Conn

	return nil
}

func(ss *SocketService) Unregister (req UnregisterReq) error {
	ss.mu.Lock()
	defer ss.mu.Unlock()

	if req.UserID == uuid.Nil {
		return ErrEmptyUserId
	}
	
	conn, ok := ss.sockets[req.UserID]

	if conn == nil {
		delete(ss.sockets, req.UserID)
		return ErrEmptyConnectionReference
	}
	
	if ok {
		conn.Close()
		delete(ss.sockets, req.UserID)
	}

	return nil
}

func (ss *SocketService) Send(
	userID uuid.UUID,
	payload any,
) (err error) {

	if userID == uuid.Nil {
		return ErrEmptyUserId
	}
	
	ss.mu.RLock()
	conn, _ := ss.sockets[userID]
	ss.mu.RUnlock()
	if conn == nil {
		return ErrEmptyConnectionReference
	}

	err = conn.WriteJSON(payload)
	if err != nil {
		return err
	}

	return nil
}