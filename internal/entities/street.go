package entities

import (
	"time"
)

type Street struct {
	ID        int       `json:"id"`
	Street    string    `json:"street"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
