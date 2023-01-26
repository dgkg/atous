package service

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"atous/db"
	"atous/model"
)

type ServiceOrder struct {
	db *db.DB
}

func initServiceOrder(r *gin.Engine, db *db.DB) {
	sa := &ServiceOrder{
		db: db,
	}

	r.POST("/order", sa.create)
	r.GET("/order/:id", sa.get)
	r.DELETE("/order/:id", sa.delete)
	r.PATCH("/order/:id", sa.update)
}

func (sa *ServiceOrder) create(c *gin.Context) {
	var order model.Order
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	order = *model.NewOrder(order.RestaurantID, order.CustomerID, order.DriverID, order.MenuID, order.Price)

	err := sa.db.CreateOrder(&order)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"order": order})
}

func (sa *ServiceOrder) get(c *gin.Context) {
	id := c.Param("id")
	order, err := sa.db.GetOrder(id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"order": order})
}

func (sa *ServiceOrder) delete(c *gin.Context) {
	id := c.Param("id")
	err := sa.db.DeleteOrder(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		return
	}
	c.JSON(http.StatusAccepted, nil)
}

func (sa *ServiceOrder) update(c *gin.Context) {
	id := c.Param("id")
	order, err := sa.db.GetOrder(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		return
	}

	newOrder := map[string]interface{}{}
	if err := c.ShouldBindJSON(&newOrder); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if value, ok := newOrder["menu_id"]; ok {
		if v, ok := value.(string); ok {
			order.MenuID = v
		}
	}
	if value, ok := newOrder["price"]; ok {
		if v, ok := value.(string); ok {
			order.Price = v
		}
	}

	err = sa.db.UpdateOrder(id, order)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"order": order})
}
