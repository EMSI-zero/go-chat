package socket

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

type Client struct {
	ID     uuid.UUID
	UserID int64
	conn   *websocket.Conn
	send   chan string
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

		//implement message logic
		fmt.Print(message)

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
		err := client.conn.WriteMessage(websocket.TextMessage, []byte(message))
		if err != nil {
			break
		}
	}
}

func NewClientSocket(hub *Hub, ws *websocket.Conn) {
	//Generate and Accept New Web socket by validating the ticket
	client := &Client{
		ID:   uuid.New(),
		conn: ws,
		send: make(chan string),
	}

	hub.register <- client

	go client.read(hub)
	go client.write()
}
