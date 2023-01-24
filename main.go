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
	r.Run()
}

func createUser(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user = *model.NewUser(user.FirstName, user.LastName, user.Age)
	c.JSON(http.StatusOK, gin.H{"user": user})
}
