package socket

import (
	"sync"

	"github.com/google/uuid"

	"github.com/gorilla/websocket"
)

type Client struct {
	ID     uuid.UUID
	UserID int64
	conn   *websocket.Conn
	send   chan string
}

type Hub struct {
	clients     map[int64]map[uuid.UUID]*Client
	broadcast   chan string
	register    chan *Client
	unregister  chan *Client
	clientsLock sync.Mutex
}

func NewHub() *Hub {
	return &Hub{
		clients:    make(map[int64](map[uuid.UUID]*Client)),
		broadcast:  make(chan string),
		register:   make(chan *Client),
		unregister: make(chan *Client),
	}
}

func (hub *Hub) InitSocketHub() {
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

func (client *Client) read(hub *Hub) {
	defer func() {
		hub.unregister <- client
		client.conn.Close()
	}()

	for {
		t, message, err := client.conn.ReadMessage()
		if err != nil {
			break
		}

		if t != websocket.TextMessage {
			break
		}
	}
}

func (client *Client) write() {
	defer func() {
		client.conn.Close()
	}()

	for message := range client.send {
		err := websocket.Message.Send(client.conn, message)
		if err != nil {
			break
		}
	}
}

func handleWebSocket(hub *Hub, ws *websocket.Conn) {
	client := &Client{
		ID: uuid.NewSHA1(),
		conn: ws,
		send: make(chan string),
	}

	hub.register <- client

	go client.read(hub)
	go client.write()
}
