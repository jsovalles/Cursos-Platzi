package websocket

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"go.uber.org/fx"
	"log"
	"net/http"
	"sync"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type Hub interface {
	HandlerWebSocket(w http.ResponseWriter, r *http.Request)
	Run()
	onConnect(client *Client)
	onDisconnect(client *Client)
	Broadcast(message any, ignore *Client)
}

type hub struct {
	clients    []*Client
	register   chan *Client
	unregister chan *Client
	mutex      *sync.Mutex
}

func NewHub() Hub {
	return &hub{
		clients:    make([]*Client, 0),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		mutex:      &sync.Mutex{},
	}
}

func (h *hub) HandlerWebSocket(w http.ResponseWriter, r *http.Request) {
	socket, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}
	client := NewClient(h, socket)
	h.register <- client

	go client.WriteMessage()
}

func (h *hub) Run() {
	for {
		select {
		case client := <-h.register:
			h.onConnect(client)
		case client := <-h.unregister:
			h.onDisconnect(client)
		}
	}
}

func (h *hub) onConnect(client *Client) {
	log.Println("Client Connected", client.socket.RemoteAddr())
	h.mutex.Lock()
	defer h.mutex.Unlock()
	client.id = client.socket.RemoteAddr().String()
	h.clients = append(h.clients, client)
}

func (h *hub) onDisconnect(client *Client) {
	log.Println("Client disconnected", client.socket.RemoteAddr())
	client.socket.Close()
	h.mutex.Lock()
	defer h.mutex.Unlock()
	i := -1
	for j, c := range h.clients {
		if c.id == client.id {
			i = j
			break
		}
	}
	copy(h.clients[i:], h.clients[i+1:])
	h.clients[len(h.clients)-1] = nil
	h.clients = h.clients[:len(h.clients)-1]
}

func (h *hub) Broadcast(message any, ignore *Client) {
	data, _ := json.Marshal(message)
	for _, client := range h.clients {
		if client != ignore {
			client.outbound <- data
		}
	}
}

var Module = fx.Provide(NewHub)
