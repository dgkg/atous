package main

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"atous/config"
	"atous/db"
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
	service.New(r, dbConn, conf.GoogleAPIKey)
	r.Run()
}
