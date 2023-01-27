package model

type Restaurant struct {
	DBData
	Name        string `json:"name"`
	Description string `json:"description"`
	Phone       string `json:"phone"`
	Email       string `json:"email"`

	Address *Address `json:"address"`

	Score string `json:"Score"`
}
