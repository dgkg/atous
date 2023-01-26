package service

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"atous/db"
	"atous/model"
)

type ServiceRestaurant struct {
	db              *db.DB
	apiGoogleMapKey string
}

func initServiceRestaurant(r *gin.Engine, db *db.DB, googleAPIKey string) {
	sr := &ServiceRestaurant{
		db:              db,
		apiGoogleMapKey: googleAPIKey,
	}
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
	restaurant, err := sr.db.GetRestaurant(id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Restaurant not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"restaurants": restaurant})
}

// deletes the user from the request body
func (sr *ServiceRestaurant) delete(c *gin.Context) {
	id := c.Param("id")
	_, err := sr.db.GetRestaurant(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Restaurant not found"})
		return
	}

	err = sr.db.DeleteRestaurant(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Restaurant not found"})
		return
	}
	c.JSON(http.StatusAccepted, nil)
}

// updates the user from the request body
func (sr *ServiceRestaurant) update(c *gin.Context) {
	id := c.Param("id")
	restaurant, err := sr.db.GetRestaurant(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Restaurant not found"})
		return
	}

	newRestaurant := map[string]interface{}{}
	if err := c.ShouldBindJSON(&newRestaurant); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if value, ok := newRestaurant["name"]; ok {
		if v, ok := value.(string); ok {
			restaurant.Name = v
		}
	}

	c.JSON(http.StatusOK, gin.H{"restaurant": restaurant})
}

func (sr *ServiceRestaurant) create(c *gin.Context) {
	var restaurant model.Restaurant
	if err := c.ShouldBindJSON(&restaurant); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := sr.db.CreateRestaurant(&restaurant)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"restaurant": restaurant})
}
