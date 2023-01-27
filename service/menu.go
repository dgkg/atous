package service

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"atous/db"
	"atous/model"
)

// ServiceMenu is the service for the menu.
type ServiceMenu struct {
	db *db.DB
}

// initServiceMenu is the initializer for the menu endpoints.
func initServiceMenu(r *gin.Engine, db *db.DB) {
	// create the service
	sm := &ServiceMenu{db: db}
	// register the endpoints
	r.POST("/menus", sm.create)
	r.GET("/menus", sm.getList)
	r.GET("/menus/:id", sm.get)
	r.DELETE("/menus/:id", sm.delete)
	r.PATCH("/menus/:id", sm.update)
}

// create is the endpoint for creating a menu.
// - It takes a JSON payload with the menu.
// - It saves the menu in the database.
// - It returns the created menu.
func (sm *ServiceMenu) create(c *gin.Context) {
	// bind the JSON payload to the model
	var payload model.Menu
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// save the menu in the database
	err := sm.db.CreateMenu(&payload)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// return the created menu
	c.JSON(http.StatusOK, gin.H{"menu": payload})
}

// getList is the endpoint for getting the list of menus.
// - It retrieves the list of menus from the database.
// - It returns the list of menus.
func (sm *ServiceMenu) getList(c *gin.Context) {
	// get the list of menus from the database
	ms, err := sm.db.GetListMenu()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// return the list of menus
	c.JSON(http.StatusOK, gin.H{"menus": ms})
}

// get is the endpoint for getting a menu by a given ID.
// - It takes the ID from the URL.
// - It retrieves the menu from the database.
// - It returns the menu.
func (sm *ServiceMenu) get(c *gin.Context) {
	// get the ID from the URL
	id := c.Param("id")
	// get the menu from the database
	menu, err := sm.db.GetMenu(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Menu not found"})
		return
	}
	// return the menu
	c.JSON(http.StatusOK, gin.H{"menu": menu})
}

// delete is the endpoint for deleting a menu by a given ID.
// - It takes the ID from the URL.
// - It deletes the menu from the database.
// - It returns an empty response with a status code 202.
func (sm *ServiceMenu) delete(c *gin.Context) {
	// get the ID from the URL
	id := c.Param("id")
	// delete the menu from the database
	err := sm.db.DeleteMenu(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Menu not found"})
		return
	}
	// return an empty response with a status code 202
	c.JSON(http.StatusAccepted, nil)
}

// update is the endpoint for updating a menu by a given ID.
// - It takes the ID from the URL.
// - It takes a JSON payload with the menu.
// - It updates the menu in the database.
// - It returns the updated menu.
func (sm *ServiceMenu) update(c *gin.Context) {
	// get the ID from the URL
	id := c.Param("id")
	// get the menu from the database
	menu, err := sm.db.GetMenu(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Menu not found"})
		return
	}
	// bind the JSON payload to a map
	payload := map[string]interface{}{}
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// update the menu fields with the payload
	setStringFromMap(payload, "name", &menu.Name)
	setIntFromMap(payload, "price", &menu.Price)
	setStringFromMap(payload, "description", &menu.Description)
	// update the menu in the database
	err = sm.db.UpdateMenu(id, menu)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// return the updated menu
	c.JSON(http.StatusOK, gin.H{"menu": menu})
}
