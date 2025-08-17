package service

import (
	"context"
	pb "github.com/go-gin-gorm-protobuf/proto"
	"github.com/google/uuid"
)

var ServerIns pb.UserServiceServer = Server{}
var persons map[string]*pb.GetUserResponse = make(map[string]*pb.GetUserResponse)

type Server struct {
	pb.UnimplementedUserServiceServer
}

func (s Server) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	person := persons[req.UserId]
	return person, nil
}

func (s Server) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	userID := uuid.NewString()
	persons[userID] = &pb.GetUserResponse{
		UserId: userID,
		Name:   req.Name,
		Age:    req.Age,
	}
	return &pb.CreateUserResponse{UserId: userID}, nil
}
