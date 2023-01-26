package service

import (
	"atous/db"
	"atous/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ServiceOrders struct {
	db *db.DB
}

func initServiceOrder(r *gin.Engine, db *db.DB) {
	sr := &ServiceOrders{db: db}
	r.POST("/orders", sr.create)
	r.GET("/orders", sr.getList)
	r.DELETE("/orders/:id", sr.delete)
}

func (sr *ServiceOrders) getList(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"orders": db.UserList})
}

// deletes the user from the request body
func (sr *ServiceOrders) delete(c *gin.Context) {
	id := c.Param("id")
	_, ok := db.UserList[id]
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		return
	}
	delete(db.UserList, id)
	c.JSON(http.StatusAccepted, nil)
}

func (sr *ServiceOrders) create(c *gin.Context) {
	var order model.Order
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	order = *model.NewOrder(order.RestaurantID, order.CustomerID)

	err := sr.db.CreateOrder(&order)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"order": order})
}
