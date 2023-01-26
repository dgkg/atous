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

	totalPrice := 0
	for k := range order.Menus {
		totalPrice += order.Menus[k].Price
	}
	order.Price = totalPrice

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

	setIntFromMap(newOrder, "price", &order.Price)

	err = sa.db.UpdateOrder(id, order)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"order": order})
}
