package services

import (
	"context"
	"errors"

	"github.com/javohir01/eater-service/src/application/dtos/address"
	"github.com/javohir01/eater-service/src/domain/address/models"
	addresssvc "github.com/javohir01/eater-service/src/domain/address/services"
)

type AddressApplicationService interface {
	CreateAddress(ctx context.Context, eaterID string, req *address.CreateAddressRequest) (*address.CreateAddressResponse, error)
	UpdateAddress(ctx context.Context, addressID string, req *address.UpdateAddressRequest) (*address.UpdateAddressResponse, error)
	DeleteAddress(ctx context.Context, addressID string) error
	GetAddress(ctx context.Context, addressID string) (*address.AddressResponse, error)
	GetAddressesByEaterID(ctx context.Context, eaterID string) ([]*address.AddressResponse, error)
}

type addressAppSvcImpl struct {
	addressSvc addresssvc.AddressService
}

func NewAddressApplicationService(addressSvc addresssvc.AddressService) AddressApplicationService {
	return &addressAppSvcImpl{
		addressSvc: addressSvc,
	}
}

func (s *addressAppSvcImpl) CreateAddress(ctx context.Context, eaterID string, req *address.CreateAddressRequest) (*address.CreateAddressResponse, error) {
	if req == nil {
		return nil, errors.New("create address request cannot be nil")
	}

	if eaterID == "" {
		return nil, errors.New("eater ID cannot be empty")
	}

	if req.Name == "" {
		return nil, errors.New("address name cannot be empty")
	}

	if req.Longitude == 0 || req.Latitude == 0 {
		return nil, errors.New("longitude and latitude cannot be zero")
	}

	addr, err := s.addressSvc.CreateAddress(ctx, eaterID, req.Name, req.Longitude, req.Latitude)
	if err != nil {
		return nil, err
	}

	return address.NewCreateAddressResponse(addr.ID, addr.Name, addr.Location.Longitude, addr.Location.Latitude), nil
}

func (s *addressAppSvcImpl) UpdateAddress(ctx context.Context, addressID string, req *address.UpdateAddressRequest) (*address.UpdateAddressResponse, error) {
	if req == nil {
		return nil, errors.New("update address request cannot be nil")
	}

	if addressID == "" {
		return nil, errors.New("address ID cannot be empty")
	}

	if req.Name == "" {
		return nil, errors.New("address name cannot be empty")
	}

	if req.Longitude == 0 || req.Latitude == 0 {
		return nil, errors.New("longitude and latitude cannot be zero")
	}

	addr, err := s.addressSvc.UpdateAddress(ctx, addressID, req.Name, req.Longitude, req.Latitude)
	if err != nil {
		return nil, err
	}

	return &address.UpdateAddressResponse{
		AddressID: addr.ID,
		Name:      addr.Name,
		Longitude: addr.Location.Longitude,
		Latitude:  addr.Location.Latitude,
	}, nil
}

func (s *addressAppSvcImpl) DeleteAddress(ctx context.Context, addressID string) error {
	if addressID == "" {
		return errors.New("address ID cannot be empty")
	}

	err := s.addressSvc.DeleteAddress(ctx, addressID)
	return err
}

func (s *addressAppSvcImpl) GetAddress(ctx context.Context, addressID string) (*address.AddressResponse, error) {
	if addressID == "" {
		return nil, errors.New("address ID cannot be empty")
	}

	addr, err := s.addressSvc.GetAddress(ctx, addressID)
	if err != nil {
		return nil, err
	}

	if addr == nil {
		return nil, errors.New("address not found")
	}

	return mapToAddressResponse(addr), nil
}

func (s *addressAppSvcImpl) GetAddressesByEaterID(ctx context.Context, eaterID string) ([]*address.AddressResponse, error) {
	if eaterID == "" {
		return nil, errors.New("eater ID cannot be empty")
	}

	addresses, err := s.addressSvc.GetAddressesByEaterID(ctx, eaterID)
	if err != nil {
		return nil, err
	}

	var responses []*address.AddressResponse
	for _, addr := range addresses {
		responses = append(responses, mapToAddressResponse(addr))
	}

	return responses, nil
}

func mapToAddressResponse(addr *models.Address) *address.AddressResponse {
	return &address.AddressResponse{
		ID:        addr.ID,
		Name:      addr.Name,
		Longitude: addr.Location.Longitude,
		Latitude:  addr.Location.Latitude,
		CreatedAt: addr.CreatedAt,
		UpdatedAt: addr.UpdatedAt,
	}
}
