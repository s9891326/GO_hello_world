package service

import (
	"github.com/gin-gonic/gin"
	"go-gin-YT/pojo"
	"log"
	"net/http"
	"strconv"
)

var userList []pojo.User

// FindAllUsers Get User
func FindAllUsers(c *gin.Context) {
	c.JSON(http.StatusOK, userList)
}

// PostUser Post User
func PostUser(c *gin.Context) {
	user := pojo.User{}
	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userList = append(userList, user)
	c.JSON(http.StatusOK, "success")
}

// DeleteUser delete user
func DeleteUser(c *gin.Context) {
	userId, _ := strconv.Atoi(c.Param("id"))
	for i, user := range userList {
		log.Println(i, user)
		if user.Id == userId {
			userList = append(userList[:i], userList[i+1:]...)
			c.JSON(http.StatusOK, "success")
			return
		}
	}
	c.JSON(http.StatusNotFound, "Error")
}

// PutUser put user
func PutUser(c *gin.Context) {
	beforeUser := pojo.User{}
	err := c.BindJSON(&beforeUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, "Error")
		return
	}

	userId, _ := strconv.Atoi(c.Param("id"))
	for i, user := range userList {
		if user.Id == userId {
			userList[i] = beforeUser
			log.Println(userList[i])
			c.JSON(http.StatusOK, "success")
			return
		}
	}

	c.JSON(http.StatusNotFound, "Error")
}
