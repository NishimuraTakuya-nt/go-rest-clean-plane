package models

import "time"

// User represents the user model
// @Description User account information
type User struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Roles     []string  `json:"roles"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
