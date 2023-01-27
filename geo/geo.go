package geo

// Geocoder is an interface for geocoding services
// It is used to get the longitude and latitude of an address
// from an address ex: "1 rue de la paix, 75000 Paris, France".
// The longitude and latitude are returned as float64.
type Geocoder interface {
	Geocode(address string) (long float64, lat float64, err error)
}
