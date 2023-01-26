package db

import (
	"fmt"
	"log"

	bolt "go.etcd.io/bbolt"

	"atous/model"
)

// TODO: to delete
var UserList = map[string]*model.User{}

const (
	BucketUsers      = "Users"
	BucketRestaurant = "Restaurants"
	BucketAddress    = "Addresses"
	BucketMenu       = "Menus"
	BucketOrder      = "Orders"
)

var modelList = []string{
	BucketUsers, BucketRestaurant,
	BucketAddress, BucketMenu, BucketOrder,
}

type DB struct {
	conn     *bolt.DB
	userList *bolt.Bucket
}

func New(dbName string) *DB {
	db, err := bolt.Open(dbName, 0600, nil)
	if err != nil {
		log.Fatal(err)
	}

	dbConn := DB{conn: db}

	err = createModel(dbConn.conn)
	if err != nil {
		log.Fatal(err)
	}

	return &dbConn
}

func createModel(db *bolt.DB) error {
	for _, model := range modelList {
		err := db.Update(func(tx *bolt.Tx) error {
			if tx.Bucket([]byte(model)) == nil {
				_, err := tx.CreateBucket([]byte(model))
				if err != nil {
					return fmt.Errorf("create bucket: %s", err)
				}
			}
			return nil
		})
		if err != nil {
			return err
		}
	}
	return nil
}
