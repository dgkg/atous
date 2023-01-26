package db

import (
	"atous/model"
	"encoding/json"
	"fmt"
	"log"

	"github.com/boltdb/bolt"
	"github.com/muyo/sno"
)

func (s *DB) Create(m *model.Menu) error {
	return s.conn.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(BucketMenu))

		m.ID = "me_" + sno.New(byte(1)).String()

		buf, err := json.Marshal(m)
		if err != nil {
			return err
		}

		return b.Put([]byte(m.ID), buf)
	})
}

func (s *DB) Update(id string, m *model.Menu) error {
	return s.conn.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(BucketMenu))

		buf, err := json.Marshal(m)
		if err != nil {
			return err
		}

		return b.Put([]byte(id), buf)
	})
}

func (s *DB) Get(id string) (*model.Menu, error) {
	var m model.Menu

	err := s.conn.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(BucketMenu))
		log.Println("GetMenu id:", id)
		v := b.Get([]byte(id))
		if v == nil {
			return fmt.Errorf("Menu not found")
		}
		return json.Unmarshal(v, &m)
	})

	if err != nil {
		return nil, err
	}

	return &m, nil
}

func (s *DB) Delete(id string) error {
	return s.conn.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(BucketMenu))
		return b.Delete([]byte(id))
	})
}

func (s *DB) GetList() ([]*model.Menu, error) {
	var menus []*model.Menu
	err := s.conn.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(BucketMenu))
		c := b.Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			var m model.Menu
			err := json.Unmarshal(v, &m)
			if err != nil {
				return err
			}
			menus = append(menus, &m)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return menus, nil
}
