package service

import (
	"atous/db"

	"github.com/gin-gonic/gin"
)

func New(r *gin.Engine, db *db.DB, googleAPIKey string) {
	initServiceUser(r, db)
	initServiceMenu(r, db)
	initServiceAddress(r, db, googleAPIKey)
}
