package db

import (
	"atous/model"
	"encoding/json"
	"errors"

	"github.com/boltdb/bolt"
	"github.com/muyo/sno"
)

func (s *DB) CreateOrder(a *model.Order) error {
	return s.conn.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(BucketRestaurant))

		a.ID = "or_" + sno.New(byte(1)).String()

		buf, err := json.Marshal(a)
		if err != nil {
			return err
		}

		return b.Put([]byte(a.ID), buf)
	})
}

func (s *DB) UpdateOrder(id string, u *model.Order) error {
	return s.conn.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(BucketRestaurant))

		buf, err := json.Marshal(u)
		if err != nil {
			return err
		}

		return b.Put([]byte(id), buf)
	})
}

func (s *DB) GetOrder(id string) (*model.Order, error) {
	var a model.Order

	err := s.conn.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(BucketRestaurant))
		v := b.Get([]byte(id))
		if v == nil {
			return errors.New("order not found")
		}
		return json.Unmarshal(v, &a)
	})

	if err != nil {
		return nil, err
	}

	return &a, nil
}

func (s *DB) DeleteOrder(id string) error {
	return s.conn.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(BucketRestaurant))
		return b.Delete([]byte(id))
	})
}
