package dtos

import (
	"time"

	pb "github.com/javohir01/eater-service/src/application/protos/eater"
	"github.com/javohir01/eater-service/src/domain/address/models"
)

func ToAddressPB(address *models.Address) *pb.Address {
	return &pb.Address{
		Id:      address.ID,
		EaterId: address.EaterID,
		Name:    address.Name,
		Location: &pb.Location{
			Latitude:  address.Location.Latitude,
			Longitude: address.Location.Longitude,
		},
		CreatedAt: address.CreatedAt.Format(time.RFC3339),
		UpdatedAt: address.UpdatedAt.Format(time.RFC3339),
	}
}

func ToAddressesPB(addresses []*models.Address) []*pb.Address {
	result := make([]*pb.Address, len(addresses))
	for i := range addresses {
		result[i] = ToAddressPB(addresses[i])
	}
	return result
}

func ValidateUpdateAddressRequestDB(req *pb.UpdateAddressRequest) error {
	return nil
}