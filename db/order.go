package db

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/muyo/sno"
	bolt "go.etcd.io/bbolt"

	"atous/model"
)

func (s *DB) CreateOrder(o *model.Order) error {
	return s.conn.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(BucketRestaurant))

		o.ID = "or_" + sno.New(byte(1)).String()
		o.CreateAt = time.Now()

		buf, err := json.Marshal(o)
		if err != nil {
			return err
		}

		return b.Put([]byte(o.ID), buf)
	})
}

func (s *DB) UpdateOrder(id string, o *model.Order) error {
	return s.conn.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(BucketRestaurant))

		o.UpdateAt = time.Now()

		buf, err := json.Marshal(o)
		if err != nil {
			return err
		}

		return b.Put([]byte(id), buf)
	})
}

func (s *DB) GetOrder(id string) (*model.Order, error) {
	var o model.Order

	err := s.conn.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(BucketRestaurant))
		v := b.Get([]byte(id))
		if v == nil {
			return errors.New("order not found")
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
		b := tx.Bucket([]byte(BucketRestaurant))
		return b.Delete([]byte(id))
	})
}
