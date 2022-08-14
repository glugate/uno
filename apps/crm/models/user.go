package models

import "time"

// User
type User struct {
	ID        int64     `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	BirthDate string    `json:"birth_date"`
	Title     string    `json:"title"`
	IsActive  bool      `json:"is_active"`
	Roles     []Role    `json:"roles"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}
