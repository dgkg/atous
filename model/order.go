package model

type Order struct {
	DBData
	RestaurantID string `json:"restaurant_id"`
	CustomerID   string `json:"customer_id"`
	DriverID     string `json:"driver_id"`
	Menus        []Menu `json:"menu_id"`
	Price        int    `json:"price"`
}
