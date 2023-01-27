package model

import "strings"

// Address represents an address and its coordinates.
type Address struct {
	DBData
	// Name of the address
	Name string `json:"name"`
	// UUID of the owner of the address, can be a restaurant or a user
	UUIDOwner string `json:"uuid_owner"`
	// Street name ex: "Rue de la Paix"
	StreetName string `json:"street_name"`
	// ZIP code number ex: "75003"
	ZIP string `json:"zip"`
	// City name ex: "Paris"
	City string `json:"city"`
	// Country name ex: "France"
	Country string `json:"country"`
	// Latitude coordinate
	Longitude float64 `json:"long"`
	// Longitude coordinate
	Latitude float64 `json:"lat"`
}

// String returns a string representation of the address in order to calculate it.
func (a *Address) String() string {
	res := a.StreetName + ", " + a.ZIP + ", " + a.City + ", " + a.Country
	res = strings.ReplaceAll(res, ", ,", ", ")
	if res[len(res)-2:] == ", " {
		res = res[:len(res)-2]
	}
	return res
}
