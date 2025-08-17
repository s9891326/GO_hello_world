package main

import (
	"context"
	"fmt"
	pb "github.com/go-gin-gorm-protobuf/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
)

func main() {
	log.Println("嘗試連接到 gRPC 伺服器...")
	// 連接 gRPC 伺服器
	conn, err := grpc.NewClient("127.0.0.1:50052", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("無法連接到伺服器: %v", err)
	}
	log.Println("成功建立連接！")
	defer conn.Close()

	client := pb.NewUserServiceClient(conn)

	// 創建用戶
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()
	createResp, err := client.CreateUser(ctx, &pb.CreateUserRequest{Name: "Alice", Age: 25})
	if err != nil {
		log.Fatalf("CreateUser failed: %v", err)
	}
	fmt.Println("Created User ID:", createResp.UserId)
	// 取得用戶
	getResp, err := client.GetUser(ctx, &pb.GetUserRequest{UserId: createResp.UserId})
	if err != nil {
		log.Fatalf("GetUser failed: %v", err)
	}
	fmt.Println("User Info:", getResp)
}
