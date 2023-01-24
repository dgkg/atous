package main

import (
	"atous/service"
	"fmt"

	"github.com/gin-gonic/gin"
)

func init() {
	fmt.Println("Init project")
}

func main() {
	r := gin.Default()
	service.New(r)
	r.Run()
}
