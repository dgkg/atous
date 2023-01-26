package model

import "github.com/google/uuid"

type Address struct {
	DBData

	UUIDOwner        string  `json:"uuid_owner"`
	StreetName       string  `json:"street_name"`
	ZIP              string  `json:"zip"`
	City             string  `json:"city"`
	GeocodeLatitude  float64 `json:"geocode_latitude"`
	GeocodeLongitude float64 `json:"geocode_longitude"`
}

func NewAddress(uuiOwner, streetname, zip, city string) *Address {
	var a Address
	a.ID = uuid.NewString()
	a.UUIDOwner = uuiOwner
	a.StreetName = streetname
	a.ZIP = zip
	a.City = city

	return &a
}
