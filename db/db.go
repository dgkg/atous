package db

import (
	"fmt"
	"log"

	bolt "go.etcd.io/bbolt"
)

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

// DB is the database struct it holds the connection to the database.
type DB struct {
	conn *bolt.DB
}

func New(dbName string) *DB {
	// Open the my.db data file in your current directory.
	db, err := bolt.Open(dbName, 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	// create the DB struct and pass the db connection
	dbConn := DB{conn: db}
	// create the buckets it they don't exist
	err = createModel(dbConn.conn)
	if err != nil {
		log.Fatal(err)
	}
	// return the DB struct
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
