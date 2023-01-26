package service

import (
	"atous/db"
	"atous/geo"

	"github.com/gin-gonic/gin"
)

func New(r *gin.Engine, db *db.DB, geocoder geo.Geocoder, googleAPIKey string) {
	initServiceUser(r, db)
	initServiceMenu(r, db)
	initServiceOrder(r, db)
	initServiceRestaurant(r, db, googleAPIKey)
	initServiceAddress(r, db, geocoder, googleAPIKey)
}
