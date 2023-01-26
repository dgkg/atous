package db

import (
	"encoding/json"
	"fmt"

	"github.com/muyo/sno"
	bolt "go.etcd.io/bbolt"

	"atous/model"
)

func (s *DB) CreateRestaurant(a *model.Restaurant) error {
	return s.conn.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(BucketRestaurant))

		a.ID = "re_" + sno.New(byte(1)).String()

		buf, err := json.Marshal(a)
		if err != nil {
			return err
		}

		return b.Put([]byte(a.ID), buf)
	})
}

func (s *DB) UpdateRestaurant(id string, u *model.Restaurant) error {
	return s.conn.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(BucketRestaurant))

		buf, err := json.Marshal(u)
		if err != nil {
			return err
		}

		return b.Put([]byte(id), buf)
	})
}

func (s *DB) GetRestaurant(id string) (*model.Restaurant, error) {
	var a model.Restaurant

	err := s.conn.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(BucketRestaurant))
		v := b.Get([]byte(id))
		if v == nil {
			return fmt.Errorf("Address not found")
		}
		return json.Unmarshal(v, &a)
	})

	if err != nil {
		return nil, err
	}

	return &a, nil
}

func (s *DB) DeleteRestaurant(id string) error {
	return s.conn.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(BucketRestaurant))
		return b.Delete([]byte(id))
	})
}

func (s *DB) GetListRestaurant() ([]*model.Restaurant, error) {
	var resp []*model.Restaurant
	err := s.conn.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(BucketRestaurant))
		c := b.Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			var u model.Restaurant
			err := json.Unmarshal(v, &u)
			if err != nil {
				return err
			}
			resp = append(resp, &u)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return resp, nil
}
