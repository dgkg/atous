package google

import (
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

// Map is a Google Map API client
// https://developers.google.com/maps/documentation/geocoding/intro
// https://developers.google.com/maps/documentation/geocoding/overview
// It's used to get the longitude and latitude of an address
// from an address ex: "1 rue de la paix, 75000 Paris, France".
// It implements the Geocoder interface.
type Map struct {
	apiKey string
}

// New returns a new Google Map API client.
// apiKey is the Google Map API key.
func New(apiKey string) *Map {
	return &Map{
		apiKey: apiKey,
	}
}

// Geocode returns the longitude and latitude of an address
// from an address ex: "1 rue de la paix, 75000 Paris, France".
// - It implements the Geocoder interface.
// - It returns an error if the address is not found.
// - It returns an error if the Google Map API returns an error.
// - It returns an error if the Google Map API returns a status different from "OK".
func (m *Map) Geocode(address string) (long float64, lat float64, err error) {
	// create the url request
	// todo: simplify this code :)
	var u url.URL
	u.Scheme = "https"
	u.Host = "maps.googleapis.com"
	u.Path = "/maps/api/geocode/json"
	value := url.Values{}
	value.Add("address", address)
	value.Add("key", m.apiKey)
	u.RawQuery = value.Encode()
	// create the request context
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	// create the request
	r, err := http.NewRequestWithContext(ctx, "GET", u.String(), nil)
	if err != nil {
		return 0, 0, err
	}
	// send the request
	resp, err := http.DefaultClient.Do(r)
	if err != nil {
		return 0, 0, err
	}
	defer resp.Body.Close()
	// read the response
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, 0, err
	}
	// unmarshal the response
	var respGoogleMap googleMapResponse
	err = json.Unmarshal(data, &respGoogleMap)
	if err != nil {
		return 0, 0, err
	}
	// check the response
	if respGoogleMap.Results == nil || respGoogleMap.Status != "OK" {
		return 0, 0, errors.New("no result")
	}
	// return the result
	return respGoogleMap.Results[0].Geometry.Location.Lng, respGoogleMap.Results[0].Geometry.Location.Lat, nil
}

// googleMapResponse is the response of the Google Map API.
// It is used to unmarshal the response.
type googleMapResponse struct {
	Results []struct {
		Geometry struct {
			Location struct {
				Lat float64 `json:"lat"`
				Lng float64 `json:"lng"`
			} `json:"location"`
		} `json:"geometry"`
	} `json:"results"`
	Status string `json:"status"`
}
