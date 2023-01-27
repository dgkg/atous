package service

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"atous/db"
	"atous/geo"
	"atous/model"
)

// ServiceAddress is the service for the address entity endpoints.
type ServiceAddress struct {
	db  *db.DB
	geo geo.Geocoder
}

// initServiceAddress initializes the address endpoints.
func initServiceAddress(r *gin.Engine, db *db.DB, geocoder geo.Geocoder) {
	// create the service
	sa := &ServiceAddress{
		db:  db,
		geo: geocoder,
	}
	// register the endpoints
	r.POST("/addresses", sa.create)
	r.GET("/addresses", sa.getList)
	r.GET("/addresses/:id", sa.get)
	r.GET("/addresses/owner/:id_owner", sa.getByOwner)
	r.DELETE("/addresses/:id", sa.delete)
	r.PATCH("/addresses/:id", sa.update)
}

// create is a POST endpoint to create a new address.
// - It takes a JSON payload with the address fields.
// - It geocodes the address.
// - Is saves the address in the database.
// - It returns the created address.
func (sa *ServiceAddress) create(c *gin.Context) {
	// bind the JSON payload to the model
	var payload model.Address
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// geocode the address
	long, lat, err := sa.geo.Geocode(payload.String())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	payload.Longitude = long
	payload.Latitude = lat
	// save the address in the database
	err = sa.db.CreateAddress(&payload)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// return the created address
	c.JSON(http.StatusOK, gin.H{"address": payload})
}

// get returns the address with the given id.
// - It get the id Parameter from the URL.
// - It retrieves the address from the database.
// - It returns the address.
func (sa *ServiceAddress) get(c *gin.Context) {
	// get the id parameter from the URL
	id := c.Param("id")
	// retrieve the address from the database
	address, err := sa.db.GetAddress(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Address not found"})
		return
	}
	// return the address
	c.JSON(http.StatusOK, gin.H{"address": address})
}

// getByOwner returns the address from its owner entity id.
// - It get the id Parameter from the URL.
// - It retrieves the address from the database.
// - It returns the address.
func (sa *ServiceAddress) getByOwner(c *gin.Context) {
	// get the id parameter from the URL
	id := c.Param("id_owner")
	// retrieve the address from the database
	address, err := sa.db.GetAddressByOwner(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Address not found"})
		return
	}
	// return the address
	c.JSON(http.StatusOK, gin.H{"address": address})
}

// delete is a DELETE endpoint to delete an address.
// - It get the id Parameter from the URL.
// - It deletes the address from the database.
// - It returns a 202 Accepted.
func (sa *ServiceAddress) delete(c *gin.Context) {
	// get the id parameter from the URL
	id := c.Param("id")
	// delete the address from the database
	err := sa.db.DeleteAddress(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Address not found"})
		return
	}
	// return a 202 Accepted
	c.JSON(http.StatusAccepted, nil)
}

// update is a PATCH endpoint to update an address.
// - It get the id Parameter from the URL.
// - It retrieves the address from the database.
// - It takes a JSON payload with a Map.
// - It updates the address fields with the payload.
// - It geocodes the address.
// - It saves the address in the database.
// - It returns the updated address.
func (sa *ServiceAddress) update(c *gin.Context) {
	// get the id parameter from the URL
	id := c.Param("id")
	// retrieve the address from the database
	address, err := sa.db.GetAddress(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Address not found"})
		return
	}
	// bind the JSON payload to a map
	payload := map[string]interface{}{}
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// update the address fields with the payload
	setStringFromMap(payload, "street_name", &address.StreetName)
	setStringFromMap(payload, "zip", &address.ZIP)
	setStringFromMap(payload, "city", &address.City)
	setStringFromMap(payload, "country", &address.Country)
	// geocode the address
	address.Longitude, address.Latitude, err = sa.geo.Geocode(address.String())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// save the address in the database
	err = sa.db.UpdateAddress(id, address)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// return the updated address
	c.JSON(http.StatusOK, gin.H{"address": address})
}

// getList returns the list of all addresses.
// - It gets the list of addresses from the database.
// - It returns the list of addresses.
func (sa *ServiceAddress) getList(c *gin.Context) {
	// get the list of addresses from the database
	addresses, err := sa.db.GetListAddress()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// return the list of addresses
	c.JSON(http.StatusOK, gin.H{"addresses": addresses})
}
