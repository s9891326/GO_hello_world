package services

import (
	"errors"
	"fmt"
	"github.com/go-gin-gorm-protobuf/internal/models"
	pb "github.com/go-gin-gorm-protobuf/proto"
	"gorm.io/gorm"
)

type UserService struct {
	DB *gorm.DB
}

func (s *UserService) CreateUser(name string, email string, password string) (*pb.User, error) {
	user := models.User{Name: name, Email: email}
	err := user.SetPassword(password)
	if err != nil {
		return nil, err
	}

	err = s.DB.Create(&user).Error
	if err != nil {
		return nil, err
	}

	return &pb.User{
		Id:    int64(user.ID),
		Name:  user.Name,
		Email: user.Email,
	}, nil
}

func (s *UserService) GetUsers() ([]*pb.User, error) {
	var users []*pb.User
	err := s.DB.Find(&users).Error
	if err != nil {
		return nil, err
	}

	var pbUsers []*pb.User
	for _, user := range users {
		pbUsers = append(pbUsers, &pb.User{
			Id:    user.Id,
			Name:  user.Name,
			Email: user.Email,
		})
	}
	return pbUsers, nil
}

func (s *UserService) DeleteUser(id uint) (string, error) {
	var user models.User

	result := s.DB.First(&user, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return "", fmt.Errorf("user with id %d not found", id)
		}
		return "", result.Error
	}

	r := s.DB.Delete(&user)
	if r.Error != nil {
		return "", r.Error
	}

	// 檢查是否真的刪除了一筆資料
	if r.RowsAffected == 0 {
		return "", fmt.Errorf("no record deleted")
	}

	return "success", nil
}

//func (s *UserService) UpdateUser(id uint, name string, email string) (*proto.User, error) {
//	var user models.User
//}
