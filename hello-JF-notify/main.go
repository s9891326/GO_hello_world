package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type NotifyRequest struct {
	Message string `json:"message" binding:"required"`
}

func main() {
	r := gin.Default()

	r.POST("/notify", func(c *gin.Context) {
		fmt.Println("Received notification:", c.Request.Body)
		// var req NotifyRequest
		// if err := c.ShouldBindJSON(&req); err != nil {
		// 	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		// 	return
		// }
		// 處理通知邏輯 (此處僅回傳收到的訊息)
		c.JSON(http.StatusOK, gin.H{
			"status":  "success",
			"message": "OK",
		})
	})

	r.Run(":8080")
}
