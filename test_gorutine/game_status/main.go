package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"sync"
	"time"
)

type gameState struct {
	Round   int      `json:"round"`
	Players []string `json:"players"`
	Status  string   `json:"status"`
}

type client struct {
	Conn   *websocket.Conn
	Send   chan []byte
	Cancel context.CancelFunc
}

type hub struct {
	clients    map[*client]bool
	register   chan *client
	unregister chan *client
	broadcast  chan []byte
	mu         sync.Mutex
}

func newHub() *hub {
	return &hub{
		clients:    make(map[*client]bool),
		register:   make(chan *client),
		unregister: make(chan *client),
		broadcast:  make(chan []byte),
	}
}

func (h *hub) run() {
	for {
		select {
		case client := <-h.register:
			h.mu.Lock()
			h.clients[client] = true
			h.mu.Unlock()
		case client := <-h.unregister:
			h.mu.Lock()
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.Send)
			}
			h.mu.Unlock()
		case message := <-h.broadcast:
			h.mu.Lock()
			for client := range h.clients {
				select {
				case client.Send <- message:
				default:
					delete(h.clients, client)
					close(client.Send)
				}
			}
			h.mu.Unlock()
		}
	}
}

func (c *client) writePump(ctx context.Context) {
	for {
		select {
		case msg, ok := <-c.Send:
			fmt.Println("writePump msg", string(msg))
			if !ok {
				// hub關閉
				_ = c.Conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			_ = c.Conn.WriteMessage(websocket.TextMessage, msg)
		case <-ctx.Done():
			// context 結束 => 關閉連線
			_ = c.Conn.Close()
			c.Cancel()
			return
		}
	}
}

func (c *client) ReadPump(ctx context.Context) {
	defer func() {
		_ = c.Conn.Close()
		c.Cancel()
	}()
	c.Conn.SetReadLimit(1024)
	for {
		_, message, err := c.Conn.ReadMessage()
		fmt.Printf("Message Received: %s\n", message)
		if err != nil {
			fmt.Println("Error reading from websocket:", err)
			break
		}

		c.Send <- message
	}
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func serveWs(hub *hub, c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println(err)
		return
	}

	client := &client{Conn: conn, Send: make(chan []byte, 256)}
	hub.register <- client

	ctx, cancel := context.WithCancel(context.Background())
	client.Cancel = cancel

	go client.writePump(ctx)
	go client.ReadPump(ctx)
}

func gameLoop(ctx context.Context, hub *hub) {
	round := 1
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Game loop stopped.")
			return
		default:
			state := gameState{
				Round:   round,
				Players: []string{"Alice", "Bob", "Cathy"},
				Status:  "playing",
			}
			data, _ := json.Marshal(state)
			hub.broadcast <- data

			round++
			time.Sleep(3 * time.Second) // 每3秒推一次狀態
		}
	}
}

func main() {
	hub := newHub()
	go hub.run()

	//http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
	//	serveWs(hub, w, r)
	//})
	//
	//fmt.Println("Server started at :8080")
	//log.Fatal(http.ListenAndServe(":8080", nil))

	// 啟動遊戲狀態推送
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	go gameLoop(ctx, hub)

	r := gin.Default()
	r.GET("/ws", func(c *gin.Context) {
		serveWs(hub, c)
	})
	r.Run(":8080")

}
