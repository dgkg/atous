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

	m := make(map[string]string)
	m["key"] = "value"
	spew.Dump(m)

	StatusCodes := map[int]string{
		200: "OK",
		201: "Created",
		202: "Accepted",
	}
	fmt.Println(StatusCodes[200])
	MapTP()

	// tbl := make([]uint, 0, 10)
	// tbl = []uint{1, 2, 3, 10, 20, 30}
	// lenTbl := len(tbl) * 200
	// for i := 0; i < lenTbl; i++ {
	// 	//fmt.Println(&tbl[i])
	// 	tbl = append(tbl, uint(i))
	// 	fmt.Println("len:", len(tbl), "cap:", cap(tbl))
	// }

	u := User{
		id:        1,
		FirstName: "John",
		LastName:  "Doe",
		Age:       42,
		RoleType:  Employee,
	}
	fmt.Println(u)
}

type Role int

const (
	Admin Role = iota
	Manager
	Employee
	MaxRole
)

var roleType = [MaxRole]string{
	Admin:    "role admin",
	Manager:  "role manager",
	Employee: "role employee",
}

func (r Role) String() string {
	return roleType[r]
}

type User struct {
	id        int
	FirstName string
	LastName  string
	Age       int
	RoleType  Role
}

func MapTP() {
	//	Créez et initialisez une map m avec en clée des strings et en valeur des ints
	m := make(map[string]int)
	// Créez une variable de type rune nommée letter et initialiser sa valeur à 'a'
	letter := 'a'
	// Créez une boucle for avec i comme itérateur sur 26 éléments
	for i := 0; i < 26; i++ {
		// Insérer dans m la valeur de letter après l’avoir casté en string
		// comme clée et comme valeur insérer l’incrément i de la boucle for
		m[string(letter)] = i
		// Auto-incrémentez la valeur de letter
		letter++
	}
	// Après la boucle for afficher la valeur de la clé "w"
	// dans la map m via la fonction fmt.Println( )
	fmt.Println(m["w"])
}
