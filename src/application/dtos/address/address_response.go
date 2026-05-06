package address

import "time"

type AddressResponse struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Longitude float64   `json:"longitude"`
	Latitude  float64   `json:"latitude"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
