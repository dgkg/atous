package service

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"atous/db"
	"atous/model"
)

type ServiceAddress struct {
	db *db.DB
}

func initServiceAddress(r *gin.Engine, db *db.DB) {
	sa := &ServiceAddress{db: db}

	r.POST("/address", sa.create)
	r.GET("/address/:id", sa.get)
	r.DELETE("/address/:id", sa.delete)
	r.PATCH("/address/:id", sa.update)
}

func (sa *ServiceAddress) create(c *gin.Context) {
	var address model.Address
	if err := c.ShouldBindJSON(&address); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	address = *model.NewAddress(address.UUIDOwner, address.StreetName, address.ZIP, address.City, address.GeocodeLatitude, address.GeocodeLongitude)

	err := sa.db.CreateAddress(&address)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"address": address})
}

func (sa *ServiceAddress) get(c *gin.Context) {
	id := c.Param("id")
	address, err := sa.db.GetAddress(id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Address not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"address": address})
}

func (sa *ServiceAddress) delete(c *gin.Context) {
	id := c.Param("id")
	err := sa.db.DeleteAddress(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Address not found"})
		return
	}
	c.JSON(http.StatusAccepted, nil)
}

func (sa *ServiceAddress) update(c *gin.Context) {
	id := c.Param("id")
	address, err := sa.db.GetAddress(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Address not found"})
		return
	}

	newAddress := map[string]interface{}{}
	if err := c.ShouldBindJSON(&newAddress); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if value, ok := newAddress["zip"]; ok {
		if v, ok := value.(string); ok {
			address.ZIP = v
		}
	}
	if value, ok := newAddress["city"]; ok {
		if v, ok := value.(string); ok {
			address.City = v
		}
	}

	err = sa.db.UpdateAddress(id, address)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"address": address})
}
