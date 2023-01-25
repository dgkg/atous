package db

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/boltdb/bolt"
	"github.com/muyo/sno"

	"atous/model"
)

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

func (s *DB) CreateUser(u *model.User) error {
	return s.conn.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(BucketUsers))

		u.ID = "us_" + sno.New(byte(1)).String()

		buf, err := json.Marshal(u)
		if err != nil {
			return err
		}

		return b.Put([]byte(u.ID), buf)
	})
}

func (s *DB) UpdateUser(id string, u *model.User) error {
	return s.conn.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(BucketUsers))

		buf, err := json.Marshal(u)
		if err != nil {
			return err
		}

		return b.Put([]byte(id), buf)
	})
}

func (s *DB) GetUser(id string) (*model.User, error) {
	var u model.User

	err := s.conn.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(BucketUsers))
		log.Println("GetUser id:", id)
		v := b.Get([]byte(id))
		if v == nil {
			return fmt.Errorf("User not found")
		}
		return json.Unmarshal(v, &u)
	})

	if err != nil {
		return nil, err
	}

	return &u, nil
}
