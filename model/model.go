package model

import "time"

// DBData is the base struct for all objects stored in the database.
type DBData struct {
	// ID is the unique identifier of the object.
	ID string `json:"id"`
	// CreateAt is the time when the object was created.
	CreateAt time.Time `json:"create_at"`
	// UpdateAt is the time when the object was last updated.
	UpdateAt time.Time `json:"update_at"`
}
