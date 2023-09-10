package dtos

import "github.com/javohir01/eater-service/src/domain/address/models"

type SaveAddressResponse struct {
	Address *models.Address 
}

func NewSaveAddressResponse(address *models.Address) *SaveAddressResponse {
	return &SaveAddressResponse{
		Address: address,
	}
}