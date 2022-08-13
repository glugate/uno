package models

import "time"

// Menu
type Menu struct {
	ID        int64      `json:"id"`
	CreatedAt time.Time  `json:"-"`
	UpdatedAt time.Time  `json:"-"`
	Label     string     `json:"label,omitempty"`
	Items     []MenuItem `json:"items"`
}
