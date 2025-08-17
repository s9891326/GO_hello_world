package main

import (
	"github.com/gin-gonic/gin"
	"go-gin-YT/database"
	. "go-gin-YT/src"
)

func main() {
	router := gin.Default()
	v1 := router.Group("/v1")
	AddUserRouter(v1)

	go func() {
		database.DBConnect()
	}()
	//router.GET("/ping", func(c *gin.Context) {
	//	c.JSON(http.StatusOK, gin.H{
	//		"message": "ping",
	//	})
	//})
	router.Run(":8000")
}
