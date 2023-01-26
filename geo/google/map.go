package google

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
)

type Map struct {
	apiKey string
}

func New(apiKey string) *Map {
	return &Map{
		apiKey: apiKey,
	}
}

func (m *Map) Geocode(address string) (long float64, lat float64, err error) {
	var u url.URL
	u.Scheme = "https"
	u.Host = "maps.googleapis.com"
	u.Path = "/maps/api/geocode/json"
	value := url.Values{}
	value.Add("address", address)
	value.Add("key", m.apiKey)
	u.RawQuery = value.Encode()

	r, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return 0, 0, err
	}
	resp, err := http.DefaultClient.Do(r)
	if err != nil {
		return 0, 0, err
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, 0, err
	}

	var respGoogleMap googleMapResponse
	err = json.Unmarshal(data, &respGoogleMap)
	if err != nil {
		return 0, 0, err
	}

	if respGoogleMap.Results == nil {
		return 0, 0, errors.New("no result")
	}

	return respGoogleMap.Results[0].Geometry.Location.Lng, respGoogleMap.Results[0].Geometry.Location.Lat, nil
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
