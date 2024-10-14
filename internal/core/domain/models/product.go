package models

import "time"

// Product represents the product model
// @Description Product information
type Product struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Price     uint      `json:"price"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
