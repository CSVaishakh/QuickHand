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
) {
	ss.mu.Lock()
	defer ss.mu.Unlock()

	ss.sockets[userID] = conn
}

func(ss *SocketService) Unregister (
	userID uuid.UUID,
){
	ss.mu.Lock()
	defer ss.mu.Unlock()

	if conn, ok := ss.sockets[userID]; ok {
		conn.Close()
		delete(ss.sockets, userID)
	}
}

func (ss *SocketService) Send(
	userId uuid.UUID,
	payload any,
) (waiting bool, err error) {
	ss.mu.RLock()
	conn, ok := ss.sockets[userId]
	ss.mu.RUnlock()
	
	if !ok {
		return true, nil
	}

	err = conn.WriteJSON(payload)
	if err != nil {
		return false, err
	}

	return false, nil
}