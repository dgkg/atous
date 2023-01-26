package service

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"

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

func geocode(key string, a *model.Address) error {

	var u url.URL
	u.Scheme = "https"
	u.Host = "maps.googleapis.com"
	u.Path = "/maps/api/geocode/json"
	value := url.Values{}
	value.Add("address", a.StreetName+" "+a.ZIP+" "+a.City)
	value.Add("key", key)
	u.RawQuery = value.Encode()

	r, err := http.NewRequest("GET", u.String(), nil)
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

	var respGoogleMap googleMapResponse
	err = json.Unmarshal(data, &respGoogleMap)
	if err != nil {
		return err
	}

	if respGoogleMap.Results == nil {
		return errors.New("no result")
	}

	a.GeocodeLatitude = respGoogleMap.Results[0].Geometry.Location.Lat
	a.GeocodeLongitude = respGoogleMap.Results[0].Geometry.Location.Lng

	return nil
}

type googleMapResponse struct {
	Results []struct {
		AddressComponents []struct {
			LongName  string   `json:"long_name"`
			ShortName string   `json:"short_name"`
			Types     []string `json:"types"`
		} `json:"address_components"`
		FormattedAddress string `json:"formatted_address"`
		Geometry         struct {
			Location struct {
				Lat float64 `json:"lat"`
				Lng float64 `json:"lng"`
			} `json:"location"`
			LocationType string `json:"location_type"`
			Viewport     struct {
				Northeast struct {
					Lat float64 `json:"lat"`
					Lng float64 `json:"lng"`
				} `json:"northeast"`
				Southwest struct {
					Lat float64 `json:"lat"`
					Lng float64 `json:"lng"`
				} `json:"southwest"`
			} `json:"viewport"`
		} `json:"geometry"`
		PlaceID  string `json:"place_id"`
		PlusCode struct {
			CompoundCode string `json:"compound_code"`
			GlobalCode   string `json:"global_code"`
		} `json:"plus_code"`
		Types []string `json:"types"`
	} `json:"results"`
	Status string `json:"status"`
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
