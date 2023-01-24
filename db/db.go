package db

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/boltdb/bolt"
	"github.com/google/uuid"

	"atous/model"
)

var UserList = map[string]*model.User{}

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

	err = db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucket([]byte("Users"))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}
		dbConn.userList = b
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}

	return &dbConn
}

func (s *DB) CreateUser(u *model.User) error {
	return s.conn.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("Users"))
		u.ID = uuid.NewString()

		buf, err := json.Marshal(u)
		if err != nil {
			return err
		}

		return b.Put([]byte(u.ID), buf)
	})
}

func (s *DB) GetUser(id string) (*model.User, error) {
	var u model.User

	err := s.conn.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("Users"))
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
