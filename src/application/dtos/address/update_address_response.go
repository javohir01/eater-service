package address

type UpdateAddressResponse struct {
	AddressID string  `json:"address_id"`
	Name      string  `json:"name"`
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latitude"`
}

func NewUpdateAddressResponse(addressID, name string, longitude, latitude float64) *UpdateAddressResponse {
	return &UpdateAddressResponse{
		AddressID: addressID,
		Name:      name,
		Longitude: longitude,
		Latitude:  latitude,
	}
}
