package address

type CreateAddressResponse struct {
	AddressID string  `json:"address_id"`
	Name      string  `json:"name"`
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latitude"`
}

func NewCreateAddressResponse(addressID, name string, longitude, latitude float64) *CreateAddressResponse {
	return &CreateAddressResponse{
		AddressID: addressID,
		Name:      name,
		Longitude: longitude,
		Latitude:  latitude,
	}
}
