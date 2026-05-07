package services

import (
	"context"

	"github.com/javohir01/eater-service/src/domain/address/models"
	"github.com/javohir01/eater-service/src/domain/address/repositories"
)

type AddressService interface {
	CreateAddress(ctx context.Context, eaterID, name string, longitude, latitude float64) (*models.Address, error)
	UpdateAddress(ctx context.Context, addressID, name string, longitude, latitude float64) (*models.Address, error)
	DeleteAddress(ctx context.Context, addressID string) error
	GetAddress(ctx context.Context, addressID string) (*models.Address, error)
	GetAddressesByEaterID(ctx context.Context, eaterID string) ([]*models.Address, error)
}

type addressSvcImpl struct {
	addressRepo repositories.AddressRepository
}

func NewAddressService(addressRepo repositories.AddressRepository) AddressService {
	return &addressSvcImpl{
		addressRepo: addressRepo,
	}
}

func (s *addressSvcImpl) CreateAddress(ctx context.Context, eaterID, name string, longitude, latitude float64) (*models.Address, error) {
	address, err := s.addressRepo.CreateAddress(ctx, eaterID, name, longitude, latitude)
	return address, err
}

func (s *addressSvcImpl) UpdateAddress(ctx context.Context, addressID, name string, longitude, latitude float64) (*models.Address, error) {
	address, err := s.addressRepo.UpdateAddress(ctx, addressID, name, longitude, latitude)
	return address, err
}

func (s *addressSvcImpl) DeleteAddress(ctx context.Context, addressID string) error {
	err := s.addressRepo.DeleteAddress(ctx, addressID)
	return err
}

func (s *addressSvcImpl) GetAddress(ctx context.Context, addressID string) (*models.Address, error) {
	address, err := s.addressRepo.GetAddress(ctx, addressID)
	if err != nil {
		return nil, err
	}
	return address, err
}

func (s *addressSvcImpl) GetAddressesByEaterID(ctx context.Context, eaterID string) ([]*models.Address, error) {
	return s.addressRepo.GetAddressesByEaterID(ctx, eaterID)
}
