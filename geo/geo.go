package geo

type Geocoder interface {
	Geocode(address string) (long float64, lat float64, err error)
}
