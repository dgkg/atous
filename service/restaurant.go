package service

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"atous/db"
	"atous/geo"
	"atous/model"
)

// ServiceRestaurant is the service for the restaurant.
type ServiceRestaurant struct {
	db  *db.DB
	geo geo.Geocoder
}

// initServiceRestaurant initializes the restaurant endpoints.
func initServiceRestaurant(r *gin.Engine, db *db.DB, geocoder geo.Geocoder) {
	// create the service
	sr := &ServiceRestaurant{
		db: db,
	}
	// register the endpoints
	r.POST("/restaurants", sr.create)
	r.GET("/restaurants", sr.getList)
	r.GET("/restaurants/:id", sr.get)
	r.DELETE("/restaurants/:id", sr.delete)
	r.PATCH("/restaurants/:id", sr.update)
}

// getList is a GET endpoint to get all the restaurants.
func (sr *ServiceRestaurant) getList(c *gin.Context) {
	// get the list of restaurants from the database
	rest, err := sr.db.GetListRestaurant()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "restaurant not found"})
		return
	}
	// return the list of restaurants
	c.JSON(http.StatusOK, gin.H{"restaurants": rest})
}

// get is a GET endpoint to get a restaurant by its id.
func (sr *ServiceRestaurant) get(c *gin.Context) {
	// get the id from the URL
	id := c.Param("id")
	// get the restaurant from the database
	restaurant, err := sr.db.GetRestaurant(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Restaurant not found"})
		return
	}
	// return the restaurant
	c.JSON(http.StatusOK, gin.H{"restaurants": restaurant})
}

// deletes the user from the request body
// - It checks if the restaurant exists into the database
// - It deletes the restaurant from the database
// - It returns a 202 Accepted with no body
func (sr *ServiceRestaurant) delete(c *gin.Context) {
	// get the id from the URL
	id := c.Param("id")
	// check if the restaurant exists into the database
	_, err := sr.db.GetRestaurant(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Restaurant not found"})
		return
	}
	// delete the restaurant from the database
	err = sr.db.DeleteRestaurant(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Restaurant not found"})
		return
	}
	// return a 202 Accepted with no body
	c.JSON(http.StatusAccepted, nil)
}

// update is a PATCH from the request body of the given restaurant ID.
// - It get the URL parameter id
// - It get the restaurant from the database
// - It updates the restaurant with the request body
// - It checks if the restaurant exists into the database
func (sr *ServiceRestaurant) update(c *gin.Context) {
	// get the id from the URL
	id := c.Param("id")
	// get the restaurant from the database
	restaurant, err := sr.db.GetRestaurant(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Restaurant not found"})
		return
	}
	// get the body of the request and bind it to a map
	newRestaurant := map[string]interface{}{}
	if err := c.ShouldBindJSON(&newRestaurant); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// update some fields of the restaurant
	setStringFromMap(newRestaurant, "name", &restaurant.Name)
	setStringFromMap(newRestaurant, "description", &restaurant.Description)
	// update the restaurant into the database
	err = sr.db.UpdateRestaurant(id, restaurant)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// return the restaurant
	c.JSON(http.StatusOK, gin.H{"restaurant": restaurant})
}

// create is a POST endpoint to create a restaurant.
// - It binds the request body to the restaurant model
// - It creates the restaurant into the database
// - It creates the address into the database
// - It returns the restaurant
func (sr *ServiceRestaurant) create(c *gin.Context) {
	// bind the request body to the restaurant model
	var restaurant model.Restaurant
	if err := c.ShouldBindJSON(&restaurant); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// create the restaurant into the database
	err := sr.db.CreateRestaurant(&restaurant)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// if the restaurant has an address, create it into the database
	if restaurant.Address != nil {
		restaurant.Address.UUIDOwner = restaurant.ID
		// if the address is not geocoded, geocode it
		// in this case, we ignore the error rerturn by the geocoder.
		restaurant.Address.Longitude, restaurant.Address.Latitude, _ = sr.geo.Geocode(restaurant.Address.String())
		err = sr.db.CreateAddress(restaurant.Address)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}
	// return the restaurant
	c.JSON(http.StatusOK, gin.H{"restaurant": restaurant})
}
