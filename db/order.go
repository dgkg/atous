package db

import (
	"atous/model"
	"encoding/json"
	"fmt"
	"log"

	"github.com/boltdb/bolt"
	"github.com/muyo/sno"
)

func (s *DB) CreateOrder(o *model.Order) error {
	return s.conn.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(BucketOrder))

		o.ID = "us_" + sno.New(byte(1)).String()

		buf, err := json.Marshal(o)
		if err != nil {
			return err
		}

		return b.Put([]byte(o.ID), buf)
	})
}

func (s *DB) GetOrders(id string) (*model.Order, error) {
	var o model.Order

	err := s.conn.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(BucketOrder))
		log.Println("GetOrder id:", id)
		v := b.Get([]byte(id))
		if v == nil {
			return fmt.Errorf("Order not found")
		}
		return json.Unmarshal(v, &o)
	})

	if err != nil {
		return nil, err
	}

	return &o, nil
}

func (s *DB) DeleteOrder(id string) error {
	return s.conn.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(BucketOrder))
		return b.Delete([]byte(id))
	})
}
