package service

import (
	"github.com/gin-gonic/gin"

	"atous/db"
	"atous/geo"
)

func New(r *gin.Engine, db *db.DB, geocoder geo.Geocoder, jwtKeySign []byte, googleAPIKey string) {
	initServiceUser(r, db, geocoder, jwtKeySign)
	initServiceMenu(r, db)
	initServiceOrder(r, db)
	initServiceRestaurant(r, db, geocoder, googleAPIKey)
	initServiceAddress(r, db, geocoder, googleAPIKey)
}
