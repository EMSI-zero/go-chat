package socket

import (
	"context"
	"fmt"

	"github.com/EMSI-zero/go-chat/domain/user"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

type Client struct {
	ID     uuid.UUID
	UserID int64
	conn   *websocket.Conn
	send   chan string
}

func (client *Client) read(ctx context.Context, hub *Hub) {
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

func NewClientSocket(ctx context.Context, hub *Hub, ws *websocket.Conn) {
	//Generate and Accept New Web socket by validating the ticket

	userId, err := user.GetUserFromCtx(ctx)
	if err != nil {
		return
	}

	client := &Client{
		ID:     uuid.New(),
		UserID: userId,
		conn:   ws,
		send:   make(chan string),
	}

	hub.register <- client

	go client.read(ctx, hub)
	go client.write()
}
