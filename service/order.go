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

	r.POST("/orders", sa.create)
	r.POST("/orders/:id/menu/:id_menu", sa.addMenu)
	r.DELETE("/orders/:id/menu/:id_menu", sa.delMenu)
	r.GET("/orders/:id", sa.get)
	r.DELETE("/orders/:id", sa.delete)
	r.PATCH("/orders/:id", sa.update)
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

func (sa *ServiceOrder) addMenu(c *gin.Context) {
	id := c.Param("id")
	order, err := sa.db.GetOrder(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "order not found"})
		return
	}

	idMenu := c.Param("id_menu")
	menu, err := sa.db.GetMenu(idMenu)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "menu not found"})
		return
	}

	// add menu to the order
	order.Menus = append(order.Menus, *menu)
	// reset price
	order.Price = 0
	for k := range order.Menus {
		order.Price += order.Menus[k].Price
	}

	err = sa.db.UpdateOrder(id, order)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"order": order})
}

func (sa *ServiceOrder) delMenu(c *gin.Context) {
	id := c.Param("id")
	order, err := sa.db.GetOrder(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		return
	}

	idMenu := c.Param("id_menu")
	for i := 0; i < len(order.Menus); i++ {
		if order.Menus[i].ID == idMenu {
			order.Menus = append(order.Menus[:i], order.Menus[i+1:]...)
			break
		}
	}

	// reset price
	order.Price = 0
	for k := range order.Menus {
		order.Price += order.Menus[k].Price
	}

	err = sa.db.UpdateOrder(id, order)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"order": order})
}
