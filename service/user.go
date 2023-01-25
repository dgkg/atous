package service

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"atous/db"
	"atous/model"
)

type ServiceUser struct {
	db *db.DB
}

func initServiceUser(r *gin.Engine, db *db.DB) {
	su := &ServiceUser{db: db}
	r.POST("/users", su.create)
	r.GET("/users", su.getList)
	r.GET("/users/:id/say-hi", su.sayHi)
	r.GET("/users/:id", su.get)
	r.DELETE("/users/:id", su.delete)
	r.PATCH("/users/:id", su.update)
}

func (su *ServiceUser) getList(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"users": db.UserList})
}

// restrives the user from the request body
func (su *ServiceUser) get(c *gin.Context) {
	id := c.Param("id")
	//user, ok := db.UserList[id]
	user, err := su.db.GetUser(id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"user": user})
}

// deletes the user from the request body
func (su *ServiceUser) delete(c *gin.Context) {
	id := c.Param("id")
	_, ok := db.UserList[id]
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	delete(db.UserList, id)
	c.JSON(http.StatusAccepted, nil)
}

// updates the user from the request body
func (su *ServiceUser) update(c *gin.Context) {
	id := c.Param("id")
	user, err := su.db.GetUser(id)
	if err != nil {
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

	if value, ok := newUser["role_type"]; ok {
		if v, ok := value.(string); ok {
			user.RoleType = model.ToRoleType(v)
		}
	}

	err = su.db.UpdateUser(id, user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
}

func (su *ServiceUser) create(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user = *model.NewUser(user.FirstName, user.LastName, &model.ConfigUser{
		Age: user.Age,
	})

	//db.UserList[user.ID] = &user

	err := su.db.CreateUser(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
}

func (su *ServiceUser) sayHi(c *gin.Context) {
	id := c.Param("id")
	user, ok := db.UserList[id]
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": user.SayHi()})
}
