package model

import (
	"time"

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

func NewOrder(restaurantID, customerID, driverID, menuID, price string) *Order {
	return &Order{
		DBData: DBData{
			ID:       uuid.NewString(),
			CreateAt: time.Now(),
		},
		RestaurantID: restaurantID,
		CustomerID:   customerID,
		DriverID:     driverID,
		MenuID:       menuID,
		Price:        price,
	}
}
