package address

type UpdateAddressRequest struct {
	Name      string  `json:"name" binding:"required"`
	Longitude float64 `json:"longitude" binding:"required"`
	Latitude  float64 `json:"latitude" binding:"required"`
}
