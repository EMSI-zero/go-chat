package socket

import (
	"sync"

	"github.com/google/uuid"
)

type Hub struct {
	clients     map[int64]map[uuid.UUID]*Client
	broadcast   chan string
	register    chan *Client
	unregister  chan *Client
	clientsLock sync.Mutex
	//dviceService
}

func NewHub() *Hub {
	return &Hub{
		clients:    make(map[int64](map[uuid.UUID]*Client)),
		broadcast:  make(chan string),
		register:   make(chan *Client),
		unregister: make(chan *Client),
	}
}

func (hub *Hub) Run() {
	for {
		select {
		case client := <-hub.register:
			hub.clientsLock.Lock()
			hub.clients[client.UserID][client.ID] = client
			hub.clientsLock.Unlock()

		case client := <-hub.unregister:
			hub.clientsLock.Lock()
			delete(hub.clients[client.UserID], client.ID)
			close(client.send)
			hub.clientsLock.Unlock()

		case message := <-hub.broadcast:
			hub.clientsLock.Lock()
			for _, clientUser := range hub.clients {
				for _, client := range clientUser {
					select {
					case client.send <- message:
					default:
						close(client.send)
						delete(hub.clients[client.UserID], client.ID)
					}
				}
			}
			hub.clientsLock.Unlock()
		}
	}
}
