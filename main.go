package main

import (
	"atous/model"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func init() {
	fmt.Println("Init project")
}

func main() {
	r := gin.Default()
	r.POST("/users", createUser)
	r.GET("/users/:id/say-hi", sayHiUser)
	r.Run()
}

func createUser(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user = *model.NewUser(user.FirstName, user.LastName, user.Age)
	model.UserList[user.ID] = &user
	c.JSON(http.StatusOK, gin.H{"user": user})
}

func sayHiUser(c *gin.Context) {
	id := c.Param("id")
	user, ok := model.UserList[id]
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": user.SayHi()})
}
