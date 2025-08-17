package pojo

import "go-gin-YT/database"

type User struct {
	Id       int    `json:"UserId"`
	Name     string `json:"UserName"`
	Password string `json:"UserPassword"`
	Email    string `json:"UserEmail"`
}

func FindAllUsers() []User {
	var users []User
	database.DB.Find(&users)
	return users
}

func FindUserById(id int) User {
	var user User
	//database.DB.First(&user, id)
	database.DB.Where("id = ?", id).First(&user)
	return user
}
