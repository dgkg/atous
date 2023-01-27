package google

import (
	"testing"
)

func TestGeocode(t *testing.T) {
	var a string
	a = "22 Rue de la Paix,"
	a += "75002,"
	a += "Paris"
	long, lat, err := New("AIzaSyAU9ZJtU14RM2QNndxY0Z8TJ2zXLwt3Fnk").Geocode(a)
	if err != nil {
		t.Error("got error", err)
	}

	if long == 0 || lat == 0 {
		t.Error("got empty geocode")
	}

	longExpected, latExpected := 2.3319534, 48.8697792
	if long != longExpected || lat != latExpected {
		t.Error("got wrong geocode", long, lat, longExpected, latExpected)
	}

}
