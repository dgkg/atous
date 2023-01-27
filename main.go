package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"atous/config"
	"atous/db"
	"atous/geo/google"
	"atous/middleware"
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
	srv := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}
	r.Use(middleware.Timeout(2 * time.Second))

	// create the database
	dbConn := db.New(conf.DBName)
	// create the geocoder
	geocoder := google.New(conf.GoogleAPIKey)
	// create the services
	service.New(r, dbConn, geocoder, conf.JWTKeySign)
	// run the server
	//r.Run()
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("listen: %s\n", err)
	}
}
