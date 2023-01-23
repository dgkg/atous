package main

import (
	"fmt"

	"github.com/davecgh/go-spew/spew"

	"atous/config"
)

func init() {
	fmt.Println("Init")
}

func main() {
	fmt.Println("Hello World")
	spew.Dump("Hello World")

	config.Exec()
	spew.Dump(config.Login)
	spew.Dump(config.Pass)
}
