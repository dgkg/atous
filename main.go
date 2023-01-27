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
	// init the config
	conf = config.New()
	fmt.Println("Init successful")
}

func main() {
	// create gin engine (router)
	r := gin.Default()
	// create the database
	dbConn := db.New(conf.DBName)
	// create the geocoder
	geocoder := google.New(conf.GoogleAPIKey)
	// create the services
	service.New(r, dbConn, geocoder, conf.JWTKeySign)
	// run the server
	r.Run()
}
