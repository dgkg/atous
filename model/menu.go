package model

type Menu struct {
	DBData
	RestaurantID string `json:"restaurant_id"`
	Name         string `json:"name"`
	Price        string `json:"price"`
}
