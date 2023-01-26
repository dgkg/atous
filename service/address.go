package service

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"atous/db"
	"atous/model"
)

type ServiceAddress struct {
	db              *db.DB
	apiGoogleMapKey string
}

func initServiceAddress(r *gin.Engine, db *db.DB, googleAPIKey string) {
	sa := &ServiceAddress{
		db:              db,
		apiGoogleMapKey: googleAPIKey,
	}

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

	address = *model.NewAddress(address.UUIDOwner, address.StreetName, address.ZIP, address.City)

	err := geocode(sa.apiGoogleMapKey, &address)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err = sa.db.CreateAddress(&address)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"address": address})
}

const (
	APIgoogleMAPS = "https://maps.googleapis.com/maps/api/geocode/json?address="
)

func geocode(key string, a *model.Address) error {
	var url = APIgoogleMAPS + a.StreetName + "," + a.City + "+" + a.ZIP + "&key=" + key
	r, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}
	resp, err := http.DefaultClient.Do(r)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	log.Println(string(data))

	return nil
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
