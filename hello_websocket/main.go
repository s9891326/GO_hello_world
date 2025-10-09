package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	pb "hello_world/hello_websocket/proto/go"
	"net/http"
	"time"
)

type Hub struct {
	Clients    map[string]*Client
	Broadcast  chan []byte
	Register   chan *Client
	Unregister chan *Client
}

type Client struct {
	ID   string           // 可以放 member id
	Conn *websocket.Conn  // 實際的 ws connection
	Send chan *pb.Message // 要送出去的資料
	Hub  *Hub             // 反向知道自己屬於哪個 hub
}

type IncomingMessage struct {
	Cmd     string `json:"cmd"`
	Payload string `json:"payload"` // payload 是 hex 字串
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.Register:
			h.Clients[client.ID] = client
		case client := <-h.Unregister:
			if _, ok := h.Clients[client.ID]; ok {
				delete(h.Clients, client.ID)
				close(client.Send)
			}
			//case message := <-h.Broadcast:
			//	for _, client := range h.Clients {
			//		select {
			//		case client.Send <- message:
			//		default:
			//			close(client.Send)
			//			delete(h.Clients, client.ID)
			//		}
			//	}
		}
	}
}

// ReadPump 讀端（接前端來的資料）
func (c *Client) ReadPump() {
	defer func() {
		c.Hub.Unregister <- c
		c.Conn.Close()
	}()
	c.Conn.SetReadLimit(512)
	for {
		_, message, err := c.Conn.ReadMessage()
		fmt.Printf("Message Received: %s\n", message)
		if err != nil {
			fmt.Println("Error reading from websocket:", err)
			break
		}

		var input IncomingMessage
		if err := json.Unmarshal(message, &input); err != nil {
			fmt.Println("Error unmarshalling from websocket:", err)
			break
		}

		msg := &pb.Message{
			Cmd:     input.Cmd,
			Payload: []byte(input.Payload),
		}
		fmt.Println("Sending message:", msg)

		c.Send <- msg
	}
}

// WritePump 寫端（推資料出去）
func (c *Client) WritePump() {
	ticker := time.NewTicker(15 * time.Second)
	defer func() {
		ticker.Stop()
		c.Conn.Close()
	}()

	for {
		select {
		case message, ok := <-c.Send:
			fmt.Printf("Send Message: %s\n", message)
			if !ok {
				c.Conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			//c.Conn.WriteMessage(websocket.TextMessage, message)
		case <-ticker.C:
			fmt.Println("ping message ticket")
			c.Conn.WriteMessage(websocket.PingMessage, []byte{})
		}
	}
}

// 用來升級 HTTP -> WebSocket
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // 開放跨域
	},
}

func ServeWs(hub *Hub, c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	client := &Client{
		ID:   "123",
		Conn: conn,
		//Send: make(chan []byte, 256),
		Send: make(chan *pb.Message),
		Hub:  hub,
	}
	fmt.Println(client.ID)
	//hub.Register <- client

	//go client.WritePump()
	//go client.ReadPump()
}

func websocketHandler(w http.ResponseWriter, r *http.Request) {
	// 把 HTTP 連線升級成 WebSocket
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Upgrade error:", err)
		return
	}
	defer conn.Close()

	// 不斷接收訊息並回傳
	for {
		mt, message, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("Read error:", err)
			break
		}
		fmt.Printf("Received: %s\n", message)

		err = conn.WriteMessage(mt, message) // 原封不動回傳
		if err != nil {
			fmt.Println("Write error:", err)
			break
		}
	}
}

func NewHub() *Hub {
	return &Hub{
		Clients:    make(map[string]*Client),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Broadcast:  make(chan []byte),
	}
}

func main() {
	// 只需要 WebSocket server → 用原生 net/http，程式更輕、更快
	//http.HandleFunc("/ws", websocketHandler)
	//fmt.Println("Server started at :8080")
	//http.ListenAndServe(":8080", nil)

	// WebSocket + REST API 混合服務 → 用 Gin，因為更容易管理 API 與 middleware。
	hub := NewHub()
	go hub.Run()

	r := gin.Default()
	r.GET("/ws", func(c *gin.Context) {
		ServeWs(hub, c)
	})
	r.Run(":8081")
}
