package db

import (
	"log"

	"github.com/boltdb/bolt"

	"atous/model"
)

var UserList = map[string]*model.User{}

type DB struct {
	conn *bolt.DB
}

func New() *DB {
	db, err := bolt.Open("my.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	return &DB{conn: db}
}
