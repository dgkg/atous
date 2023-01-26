package model

import (
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/google/uuid"
)

type Menu struct {
	DBData
	RestaurantID string `json:"restaurant_id"`
	Name         string `json:"name"`
	Price        string `json:"price"`
	ImageURI     string `json:"image"`
}

func AddMenu(name, price, imageURI string) *Menu {
	var m Menu
	m.RestaurantID = uuid.NewString()
	m.Name = name
	m.Price = price
	m.ImageURI = imageURI
	m.CreateAt = time.Now()

	spew.Printf("Model : AddMenu: %#v\n", m)
	return &m
}
