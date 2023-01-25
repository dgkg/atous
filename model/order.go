package model

type Order struct {
	DBData
	RestaurantID string `json:"restaurant_id"`
	CustomerID   string `json:"customer_id"`
	DriverID     string `json:"driver_id"`
	MenuID       string `json:"menu_id"`
	Price        string `json:"price"`
}
