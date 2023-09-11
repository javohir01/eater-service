package services

import (
	"context"
	"errors"

	dtos "github.com/javohir01/eater-service/src/application/dtos/address"
	addressSvc "github.com/javohir01/eater-service/src/domain/address/services"
)

type AddressApplicationService interface {
	CreateAddress(ctx context.Context, addressID, EaterID, name string, Latitude, Longitude float64) (*dtos.SaveAddressResponse, error)
	UpdateAddress(ctx context.Context, addressID, EaterID, name string, Latitude, Longitude float64) (*dtos.SaveAddressResponse, error)
	DeleteAddress(ctx context.Context, addressID string) error
	GetAddressById(ctx context.Context, addressID string) (*dtos.SaveAddressResponse, error)
	ListAddressByEaterId(ctx context.Context, eaterID string) ([]*dtos.ListAddressByEaterIdResponse, error)
}

type addressAppSvcImpl struct {
	addressSvc addressSvc.AddressService
}

func NewAddressApplicationService(
	addressSvc addressSvc.AddressService,
) AddressApplicationService {
	return &addressAppSvcImpl{
		addressSvc: addressSvc,
	}
}

func (s *addressAppSvcImpl) CreateAddress(ctx context.Context, addressID, EaterID, name string, Latitude, Longitude float64) (*dtos.SaveAddressResponse, error) {
	if addressID == "" {
		return nil, errors.New("address is is required!")
	}

	if EaterID == "" {
		return nil, errors.New("EaterID is is required!")
	}

	if name == "" {
		return nil, errors.New("name is is required!")
	}

	if Latitude == 0 {
		return nil, errors.New("Latitude is is required!")
	}

	if Longitude == 0 {
		return nil, errors.New("Longitude is is required!")
	}

	result, err := s.CreateAddress(ctx, addressID, EaterID, name, Latitude, Longitude)
	if err != nil {
		return nil, err
	}

	return dtos.SaveAddressResponse(result), nil
}

func (s *addressAppSvcImpl) UpdateAddress(ctx context.Context, addressID, EaterID, name string, Latitude, Longitude float64) (*dtos.SaveAddressResponse, error) {
	if addressID == "" {
		return nil, errors.New("address is is required!")
	}

	if EaterID == "" {
		return nil, errors.New("EaterID is is required!")
	}

	if name == "" {
		return nil, errors.New("name is is required!")
	}

	if Latitude == 0 {
		return nil, errors.New("Latitude is is required!")
	}

	if Longitude == 0 {
		return nil, errors.New("Longitude is is required!")
	}

	result, err := s.UpdateAddress(ctx, addressID, EaterID, name, Latitude, Longitude)
	if err != nil {
		return nil, err
	}

	return dtos.SaveAddressResponse(result), nil
}

func (s *addressAppSvcImpl) DeleteAddress(ctx context.Context, addressID string) error {
	if addressID == "" {
		return errors.New("addressID is is required!")
	}

	err := s.DeleteAddress(ctx, addressID)
	if err != nil {
		return err
	}

	return nil
}

func (s *addressAppSvcImpl) GetAddressById(ctx context.Context, addressID string) (*dtos.SaveAddressResponse, error) {
	if addressID == "" {
		return nil, errors.New("addressID is is required!")
	}

	result, err := s.GetAddressById(ctx, addressID)
	if err != nil {
		return nil, err
	}

	response := dtos.SaveAddressResponse(result)
	return &response, nil
}

func (s *addressAppSvcImpl) ListAddressByEaterId(ctx context.Context, eaterID string) ([]*dtos.ListAddressByEaterIdResponse, error) {
	if eaterID == "" {
		return nil, errors.New("eaterID is is required!")
	}

	result, err := s.ListAddressByEaterId(ctx, eaterID)
	if err != nil {
		return nil, err
	}

	return dtos.SaveAddressResponse(result), nil
}
