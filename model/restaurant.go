package model

type Restaurant struct {
	DBData
	Name  string `json:"name"`
	Phone string `json:"phone"`
	Email string `json:"email"`

	Address string `json:"address"`
	ZIP     string `json:"zip"`
	City    string `json:"city"`

	GeocodeLatitude  string `json:"geocode_latitude"`
	GeocodeLongitude string `json:"geocode_longitude"`

	Score string `json:"Score"`
}

type Menu struct {
	DBData
	RestaurantID string `json:"restaurant_id"`
	Name         string `json:"name"`
	Price        string `json:"price"`
}

type Order struct {
	DBData
	RestaurantID string `json:"restaurant_id"`
	CustomerID   string `json:"customer_id"`
	DriverID     string `json:"driver_id"`
	MenuID       string `json:"menu_id"`
	Price        string `json:"price"`
}
