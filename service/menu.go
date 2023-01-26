package service

import (
	"atous/db"
	"atous/model"
	"net/http"

	"github.com/gin-gonic/gin"
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
	var menu model.Menu
	if err := c.ShouldBindJSON(&menu); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	menu = *model.AddMenu(menu.RestaurantID, menu.Name, menu.Price, menu.ImageURI)

	err := sm.db.CreateMenu(&menu)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"menu": menu})
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

	newMenu := map[string]interface{}{}
	if err := c.ShouldBindJSON(&newMenu); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if value, ok := newMenu["name"]; ok {
		if v, ok := value.(string); ok {
			menu.Name = v
		}
	}
	if value, ok := newMenu["price"]; ok {
		if v, ok := value.(string); ok {
			menu.Price = v
		}
	}

	err = sm.db.UpdateMenu(id, menu)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"menu": menu})
}
