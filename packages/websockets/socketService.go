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

func(ss *SocketService) Register (
	userID uuid.UUID,
	conn *websocket.Conn,
) error {

	if userID == uuid.Nil {
		return ErrEmptyUserId
	}

	if conn == nil {
		return ErrEmptyConnectionReference
	}
	
	ss.mu.Lock()
	defer ss.mu.Unlock()

	if oldConn, exists := ss.sockets[userID]; exists {
		_ = oldConn.Close()
	} 

	ss.sockets[userID] = conn

	return nil
}

func(ss *SocketService) Unregister (
	userID uuid.UUID,
) error {
	ss.mu.Lock()
	defer ss.mu.Unlock()

	if userID == uuid.Nil {
		return ErrEmptyUserId
	}

	conn, ok := ss.sockets[userID]

	if conn == nil {
		return ErrEmptyConnectionReference
	}
	
	if ok {
		conn.Close()
		delete(ss.sockets, userID)
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