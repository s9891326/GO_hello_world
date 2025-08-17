package main

import (
	"log"
	"net"

	pb "github.com/go-gin-gorm-protobuf/proto"
	"github.com/go-gin-gorm-protobuf/service"
	"google.golang.org/grpc"
)

func main() {
	// 開啟 gRPC 服務
	lis, err := net.Listen("tcp", "localhost:50052")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterUserServiceServer(s, service.ServerIns)

	log.Println("Server is running on port 50052...")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
