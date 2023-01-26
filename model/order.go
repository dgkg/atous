package model

import (
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/google/uuid"
)

type Order struct {
	DBData
	RestaurantID string `json:"restaurant_id"`
	CustomerID   string `json:"customer_id"`
	DriverID     string `json:"driver_id"`
	MenuID       string `json:"menu_id"`
	Price        string `json:"price"`
}

func NewOrder(restaurantID string, customer_id string) *Order {
	var o Order
	o.ID = uuid.NewString()
	o.RestaurantID = restaurantID
	o.CustomerID = customer_id

	o.CreateAt = time.Now()

	spew.Printf("Model : New Order: %#v\n", o)
	return &o
}
