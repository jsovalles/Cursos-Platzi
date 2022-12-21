package websocket

import "github.com/gorilla/websocket"

type Client struct {
	hub      *hub
	id       string
	socket   *websocket.Conn
	outbound chan []byte
}

func NewClient(hub *hub, socket *websocket.Conn) *Client {
	return &Client{
		hub:      hub,
		socket:   socket,
		outbound: make(chan []byte),
	}
}

func (c *Client) WriteMessage() {
	for {
		select {
		case message, ok := <-c.outbound:
			if !ok {
				c.socket.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			c.socket.WriteMessage(websocket.TextMessage, message)
		}
	}
}
