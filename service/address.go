package service

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"atous/db"
	"atous/geo"
	"atous/model"
)

type ServiceAddress struct {
	db              *db.DB
	apiGoogleMapKey string
	geo             geo.Geocoder
}

func initServiceAddress(r *gin.Engine, db *db.DB, geocoder geo.Geocoder, googleAPIKey string) {
	sa := &ServiceAddress{
		db:              db,
		apiGoogleMapKey: googleAPIKey,
		geo:             geocoder,
	}

	r.POST("/address", sa.create)
	r.GET("/address/:id", sa.get)
	r.DELETE("/address/:id", sa.delete)
	r.PATCH("/address/:id", sa.update)
}

func (sa *ServiceAddress) create(c *gin.Context) {
	var payload model.Address
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	long, lat, err := sa.geo.Geocode(payload.String())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	payload.Longitude = long
	payload.Latitude = lat

	err = sa.db.CreateAddress(&payload)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"address": payload})
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

	payload := map[string]interface{}{}
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	setStringFromMap(payload, "street_name", &address.StreetName)
	setStringFromMap(payload, "zip", &address.ZIP)
	setStringFromMap(payload, "city", &address.City)
	setStringFromMap(payload, "country", &address.Country)

	address.Longitude, address.Latitude, err = sa.geo.Geocode(address.String())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err = sa.db.UpdateAddress(id, address)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"address": address})
}
