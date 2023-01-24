package service

import (
	"atous/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func getListUsers(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"users": model.UserList})
}

// restrives the user from the request body
func getUser(c *gin.Context) {
	id := c.Param("id")
	user, ok := model.UserList[id]
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"user": user})
}

// deletes the user from the request body
func deleteUser(c *gin.Context) {
	id := c.Param("id")
	_, ok := model.UserList[id]
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	delete(model.UserList, id)
	c.JSON(http.StatusAccepted, nil)
}

// updates the user from the request body
func updateUser(c *gin.Context) {
	id := c.Param("id")
	user, ok := model.UserList[id]
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	newUser := map[string]interface{}{}
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if value, ok := newUser["first_name"]; ok {
		if v, ok := value.(string); ok {
			user.FirstName = v
		}
	}
	if value, ok := newUser["last_name"]; ok {
		if v, ok := value.(string); ok {
			user.LastName = v
		}
	}
	if value, ok := newUser["age"]; ok {
		if v, ok := value.(int); ok {
			user.Age = v
		}
	}
	c.JSON(http.StatusOK, gin.H{"user": user})
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
