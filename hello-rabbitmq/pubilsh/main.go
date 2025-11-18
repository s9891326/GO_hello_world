package main

import (
	"encoding/json"
	"log"
	"os"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

// 使用範例：AMQP_URL 可透過環境變數設定，預設連到本機 rabbitmq
// go run main.go
func main() {
	amqpURL := os.Getenv("AMQP_URL")
	if amqpURL == "" {
		amqpURL = "amqp://guest:guest@localhost:5672/"
	}

	conn, err := amqp.Dial(amqpURL)
	if err != nil {
		log.Fatalf("連線到 RabbitMQ 失敗: %v", err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("開啟 channel 失敗: %v", err)
	}
	defer ch.Close()

	// 宣告隊列（如果不存在會被建立）
	q, err := ch.QueueDeclare(
		"hello_json4", // name
		false,         // durable
		false,         // delete when unused
		false,         // exclusive
		false,         // no-wait
		nil,           // args
	)
	if err != nil {
		log.Fatalf("宣告隊列失敗: %v", err)
	}

	// 範例要傳送的 JSON 內容
	payload := struct {
		ID        string    `json:"id"`
		Message   string    `json:"message"`
		Timestamp time.Time `json:"timestamp"`
	}{
		ID:        "1",
		Message:   "Hello, RabbitMQ!",
		Timestamp: time.Now().UTC(),
	}

	body, err := json.Marshal(payload)
	if err != nil {
		log.Fatalf("序列化 JSON 失敗: %v", err)
	}

	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key = queue 名稱
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)
	if err != nil {
		log.Fatalf("發佈訊息失敗: %v", err)
	}

	log.Printf("已發佈 JSON 訊息到隊列 %s: %s", q.Name, string(body))
}
