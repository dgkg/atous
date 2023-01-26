package service

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"atous/db"
	"atous/model"
)

type ServiceMenu struct {
	db *db.DB
}

func initServiceMenu(r *gin.Engine, db *db.DB) {
	sm := &ServiceMenu{db: db}
	r.POST("/menu", sm.create)
	r.GET("/menu", sm.getList)
	r.GET("/menu/:id", sm.get)
	r.DELETE("/menu/:id", sm.delete)
	r.PATCH("/menu/:id", sm.update)
}

func (sm *ServiceMenu) create(c *gin.Context) {
	var payload model.Menu
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := sm.db.CreateMenu(&payload)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"menu": payload})
}

func (sm *ServiceMenu) getList(c *gin.Context) {
	ms, err := sm.db.GetListMenu()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"menus": ms})
}

func (sm *ServiceMenu) get(c *gin.Context) {
	id := c.Param("id")

	menu, err := sm.db.GetMenu(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Menu not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"menu": menu})
}

func (sm *ServiceMenu) delete(c *gin.Context) {
	id := c.Param("id")
	err := sm.db.DeleteMenu(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Menu not found"})
		return
	}
	c.JSON(http.StatusAccepted, nil)
}

func (sm *ServiceMenu) update(c *gin.Context) {
	id := c.Param("id")
	menu, err := sm.db.GetMenu(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Menu not found"})
		return
	}

	payload := map[string]interface{}{}
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	setStringFromMap(payload, "name", &menu.Name)
	setIntFromMap(payload, "price", &menu.Price)
	setStringFromMap(payload, "description", &menu.Description)

	err = sm.db.UpdateMenu(id, menu)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"menu": menu})
}

func setStringFromMap(m map[string]interface{}, key string, field *string) {
	if value, ok := m[key]; ok {
		if v, ok := value.(string); ok {
			*field = v
		}
	}
}

func setIntFromMap(m map[string]interface{}, key string, field *int) {
	if value, ok := m[key]; ok {
		if v, ok := value.(int); ok {
			*field = v
		}
	}
}
