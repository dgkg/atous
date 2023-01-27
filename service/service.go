package service

import (
	"github.com/gin-gonic/gin"

	"atous/db"
	"atous/geo"
)

// New initializes the services endpoints of the current entities:
// - user
// - menu
// - order
// - restaurant
// - address
// It use in params a gin engine, a database connection, a geocoder, a JWT key sign.
func New(r *gin.Engine, db *db.DB, geocoder geo.Geocoder, jwtKeySign []byte) {
	initServiceUser(r, db, geocoder, jwtKeySign)
	initServiceMenu(r, db)
	initServiceOrder(r, db)
	initServiceRestaurant(r, db, geocoder)
	initServiceAddress(r, db, geocoder)
}
