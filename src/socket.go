package main

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	ws "github.com/gorilla/websocket"
	"log"
	"net/http"
)

var upgrader = ws.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

type Pool struct {
	Register     chan *Client
	Unregister   chan *Client
	Clients      map[*Client]bool
	Broadcast    chan MessageWithClient
	BroadcastAll chan Message
	Send         chan MessageWithClient
}

func NewPool() *Pool {
	return &Pool{
		Register:     make(chan *Client),
		Unregister:   make(chan *Client),
		Clients:      make(map[*Client]bool),
		Broadcast:    make(chan MessageWithClient),
		BroadcastAll: make(chan Message),
		Send:         make(chan MessageWithClient),
	}
}

type Message struct {
	Type string            `json:"type"`
	Body map[string]string `json:"body"`
}

type Client struct {
	ID   string
	Conn *ws.Conn
	Pool *Pool
}
type MessageWithClient struct {
	Message Message
	Client  *Client
}

func (c *Client) Read() {
	defer func() {
		c.Pool.Unregister <- c
		c.Conn.Close()
	}()

	for {
		_, p, err := c.Conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		var event map[string]string
		err = json.Unmarshal(p, &event)
		if err != nil {
			log.Println(err)
			return
		}
		HandleMessage(c, event)
	}
}

func (pool *Pool) Start() {
	for {
		select {
		case client := <-pool.Register:
			pool.Clients[client] = true
			fmt.Println("New client. Pool size: ", len(pool.Clients))
			go pool.HandleRegister(client)
			break
		case client := <-pool.Unregister:
			delete(pool.Clients, client)
			go pool.HandleUnregister(client)
			break
		case messageWithClient := <-pool.Broadcast:
			if err := Broadcast(pool, messageWithClient); err != nil {
				fmt.Println(err)
				return
			}
		case message := <-pool.BroadcastAll:
			if err := BroadcastAll(pool, message); err != nil {
				fmt.Println(err)
				return
			}
		case messageWithClient := <-pool.Send:
			err := messageWithClient.Client.Conn.WriteJSON(messageWithClient.Message)
			if err != nil {
				fmt.Println(err)
				return
			}
		}
	}
}

func Broadcast(pool *Pool, m MessageWithClient) error {
	for clientItem, _ := range pool.Clients {
		if clientItem.ID != m.Client.ID {
			err := clientItem.Conn.WriteJSON(m.Message)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func BroadcastAll(pool *Pool, m Message) error {
	for clientItem, _ := range pool.Clients {
		err := clientItem.Conn.WriteJSON(m)
		if err != nil {
			return err
		}
	}
	return nil
}

func serveWs(pool *Pool, w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	client := &Client{
		ID:   uuid.NewString(),
		Conn: conn,
		Pool: pool,
	}

	pool.Register <- client
	client.Read()
}
