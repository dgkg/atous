package model

type Address struct {
	DBData

	UUIDOwner string `json:"uuid_owner"`

	StreetName       string `json:"street_name"`
	ZIP              string `json:"zip"`
	City             string `json:"city"`
	GeocodeLatitude  string `json:"geocode_latitude"`
	GeocodeLongitude string `json:"geocode_longitude"`
}
