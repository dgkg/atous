package db

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/muyo/sno"
	bolt "go.etcd.io/bbolt"

	"atous/model"
)

func (s *DB) CreateMenu(m *model.Menu) error {
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

func (s *DB) UpdateMenu(id string, m *model.Menu) error {
	return s.conn.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(BucketMenu))

		buf, err := json.Marshal(m)
		if err != nil {
			return err
		}

		return b.Put([]byte(id), buf)
	})
}

func (s *DB) GetMenu(id string) (*model.Menu, error) {
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

func (s *DB) DeleteMenu(id string) error {
	return s.conn.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(BucketMenu))
		return b.Delete([]byte(id))
	})
}

func (s *DB) GetListMenu() ([]*model.Menu, error) {
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

func (s *DB) GetMenuByRestaurant(idRestaurant string) ([]*model.Menu, error) {
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

	var menusRestaurant []*model.Menu
	for k := range menus {
		if menus[k].RestaurantID == idRestaurant {
			menusRestaurant = append(menusRestaurant, menus[k])
		}
	}
	return menusRestaurant, nil
}
