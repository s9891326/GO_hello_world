package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-gin-gorm-protobuf/config"
	"github.com/go-gin-gorm-protobuf/internal/controllers"
	"github.com/go-gin-gorm-protobuf/internal/models"
	"github.com/go-gin-gorm-protobuf/internal/services"
)

func main() {
	r := gin.Default()

	config.ConnectDatabase()
	models.Migrate(config.DB)

	userService := &services.UserService{DB: config.DB}
	userController := &controllers.UserController{Service: userService}

	r.GET("/users", userController.GetUser)
	r.POST("/users", userController.CreateUser)
	r.DELETE("/users/:id", userController.DeleteUser)
	err := r.Run(":8000")
	if err != nil {
		return
	}
}
