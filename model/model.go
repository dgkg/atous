package model

import "time"

type DBData struct {
	ID string `json:"id"`
	// DB dates.
	CreateAt time.Time `json:"create_at"`
	UpdateAt time.Time `json:"update_at"`
}
