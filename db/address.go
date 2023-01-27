package db

import (
	"encoding/json"
	"errors"
	"log"
	"time"

	"github.com/muyo/sno"
	bolt "go.etcd.io/bbolt"

	"atous/model"
)

func (s *DB) CreateAddress(a *model.Address) error {
	return s.conn.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(BucketAddress))

		a.ID = "ad_" + sno.New(byte(1)).String()
		a.CreateAt = time.Now()

		buf, err := json.Marshal(a)
		if err != nil {
			return err
		}

		return b.Put([]byte(a.ID), buf)
	})
}

func (s *DB) UpdateAddress(id string, a *model.Address) error {
	return s.conn.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(BucketAddress))

		a.UpdateAt = time.Now()

		buf, err := json.Marshal(a)
		if err != nil {
			return err
		}

		return b.Put([]byte(id), buf)
	})
}

func (s *DB) GetAddress(id string) (*model.Address, error) {
	var a model.Address

	err := s.conn.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(BucketAddress))
		v := b.Get([]byte(id))
		if v == nil {
			return errors.New("address not found")
		}
		return json.Unmarshal(v, &a)
	})

	if err != nil {
		return nil, err
	}

	return &a, nil
}

func (s *DB) DeleteAddress(id string) error {
	return s.conn.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(BucketAddress))
		return b.Delete([]byte(id))
	})
}

func (s *DB) GetListAddress() ([]*model.Address, error) {
	var adds []*model.Address
	err := s.conn.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(BucketAddress))
		c := b.Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			var a model.Address
			err := json.Unmarshal(v, &a)
			if err != nil {
				return err
			}
			adds = append(adds, &a)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return adds, nil
}

func (s *DB) GetAddressByOwner(idOwner string) (*model.Address, error) {
	log.Println("GetAddressByOwner", idOwner)
	var add *model.Address
	err := s.conn.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(BucketAddress))
		c := b.Cursor()
		var a model.Address
		for k, v := c.First(); k != nil; k, v = c.Next() {
			err := json.Unmarshal(v, &a)
			if err != nil {
				return err
			}
			log.Println("GetAddressByOwner search:", a.UUIDOwner)
			if a.UUIDOwner == idOwner {
				log.Println("GetAddressByOwner found:", add)
				add = &a
				return nil
			}
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return add, nil
}
