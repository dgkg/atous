package service

import (
	"atous/model"
	"testing"
)

func TestGeocode(t *testing.T) {
	var a model.Address
	a.StreetName = "22 Rue de la Paix"
	a.ZIP = "75002"
	a.City = "Paris"
	err := geocode("AIzaSyAU9ZJtU14RM2QNndxY0Z8TJ2zXLwt3Fnk", &a)
	if err != nil {
		t.Error("got error", err)
	}

	if a.GeocodeLatitude == 0 || a.GeocodeLongitude == 0 {
		t.Error("got empty geocode")
	}

	if a.GeocodeLatitude != 48.8697792 || a.GeocodeLongitude != 2.3319534 {
		t.Error("got wrong geocode")
	}

}
