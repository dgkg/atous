package geo

// Geocoder is an interface for geocoding services
// It is used to get the longitude and latitude of an address
// from an address ex: "1 rue de la paix, 75000 Paris, France".
// The longitude and latitude are returned as float64.
type Geocoder interface {
	// Geocode returns the longitude and latitude of an address.
	// It returns an error if the geocoding failed.
	// It returns 0, 0 if the address is not found.
	Geocode(address string) (long float64, lat float64, err error)
}
