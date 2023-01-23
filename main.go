package main

import (
	"fmt"

	"github.com/davecgh/go-spew/spew"

	"atous/a"
)

func main() {
	fmt.Println("Hello World")
	spew.Dump("Hello World")

	a.Exec()
	spew.Dump(a.Coucou)
	spew.Dump(a.Toto)
}

// package
// import
// type
// var
// const
// if
// continue
// else
// switch
// case
// fallthrough
// default
// break
// defer
// range
// return
// map
// struct
// func
// interface
// chan
// for
// go
// select

// goto
