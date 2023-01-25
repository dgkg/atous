package service

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"atous/db"
	"atous/model"
)

type ServiceRestaurant struct {
	db *db.DB
}

func initServiceRestaurant(r *gin.Engine, db *db.DB) {
	sr := &ServiceRestaurant{db: db}
	r.POST("/restaurants", sr.create)
	r.GET("/restaurants", sr.getList)
	r.GET("/restaurants/:id", sr.get)
	r.DELETE("/restaurants/:id", sr.delete)
	r.PATCH("/restaurants/:id", sr.update)
}

func (sr *ServiceRestaurant) getList(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"restaurants": db.UserList})
}

// restrives the user from the request body
func (sr *ServiceRestaurant) get(c *gin.Context) {
	id := c.Param("id")
	//user, ok := db.UserList[id]
	user, err := sr.db.GetUser(id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"restaurants": user})
}

// deletes the user from the request body
func (sr *ServiceRestaurant) delete(c *gin.Context) {
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
func (sr *ServiceRestaurant) update(c *gin.Context) {
	id := c.Param("id")
	user, ok := db.UserList[id]
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

func (sr *ServiceRestaurant) create(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user = *model.NewUser(user.FirstName, user.LastName, &model.ConfigUser{
		Age: user.Age,
	})

	//db.UserList[user.ID] = &user

	err := sr.db.CreateUser(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
}
