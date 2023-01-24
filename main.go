package main

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"atous/db"
	"atous/service"
)

func init() {
	fmt.Println("Init project")
}

func main() {
	r := gin.Default()
	dbConn := db.New()
	service.New(r, dbConn)
	r.Run()

}
