package service

import (
	"atous/db"

	"github.com/gin-gonic/gin"
)

func New(r *gin.Engine, db *db.DB, googleAPIKey string) {
	initServiceUser(r, db)
	initServiceMenu(r, db)
	initServiceOrder(r, db)
	initServiceRestaurant(r, db, googleAPIKey)
	initServiceAddress(r, db, googleAPIKey)
}
