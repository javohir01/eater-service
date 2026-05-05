package models

import "time"

type Location struct {
	Longitude float64
	Latitude  float64
}

type Address struct {
	ID        string
	EaterID   string
	Name      string
	Location  *Location
	CreatedAt time.Time
	UpdatedAt time.Time
}
