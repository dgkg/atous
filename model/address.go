package model

import "strings"

type Address struct {
	DBData

	UUIDOwner  string  `json:"uuid_owner"`
	StreetName string  `json:"street_name"`
	ZIP        string  `json:"zip"`
	City       string  `json:"city"`
	Country    string  `json:"country"`
	Latitude   float64 `json:"lat"`
	Longitude  float64 `json:"long"`
}

func (a *Address) String() string {
	res := a.StreetName + ", " + a.ZIP + ", " + a.City + ", " + a.Country
	res = strings.ReplaceAll(res, ", ,", ", ")
	if res[len(res)-2:] == ", " {
		res = res[:len(res)-2]
	}
	return res
}
