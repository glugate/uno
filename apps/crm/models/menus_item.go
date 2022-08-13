package models

import "time"

// Menus Item
type MenuItem struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
	MenuId    int64     `json:"menus_id"`
	ParentId  int64     `json:"parent_id"`
	Label     string    `json:"label"`
	Path      string    `json:"path"`
}
