package dtos

type GetAddressResponse struct {
	AddressID string `json:"address_id"`
}

func NewGetAddressResponse(addressID string) *GetAddressResponse {
	return &GetAddressResponse{
		AddressID: addressID,
	}
}
