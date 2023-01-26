package main

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"atous/config"
	"atous/db"
	"atous/geo/google"
	"atous/service"
)

var conf *config.Config

func init() {
	fmt.Println("Init project")
	conf = config.New()
	fmt.Println("Init successful")
}

func main() {
	r := gin.Default()
	dbConn := db.New(conf.DBName)
	geocoder := google.New(conf.GoogleAPIKey)
	service.New(r, dbConn, geocoder, conf.GoogleAPIKey)
	r.Run()
}
