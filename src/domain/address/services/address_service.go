package services

import (
	"context"
	"time"

	"github.com/javohir01/eater-service/src/domain/address/models"
	"github.com/javohir01/eater-service/src/domain/address/repositories"
	"github.com/javohir01/eater-service/src/infastructure/rand"
)

type AddressService interface {
	CreateAddress(ctx context.Context, EaterID, name string, Latitude, Longitude float64) (*models.Address, error)
	UpdateAddress(ctx context.Context, addressID, EaterID, name string, Latitude, Longitude float64) (*models.Address, error)
	DeleteAddress(ctx context.Context, addressID string) error
	GetAddressById(ctx context.Context, addressID string) (*models.Address, error)
	ListAddressByEaterId(ctx context.Context, eaterID string) ([]*models.Address, error)
}

type addressSvcImpl struct {
	addressRepo repositories.AddressRepository
}

func NewAddressService(
	addressRepo repositories.AddressRepository,
) AddressService {
	return &addressSvcImpl{
		addressRepo: addressRepo,
	}
}

func (s *addressSvcImpl) CreateAddress(ctx context.Context, EaterID, name string, Latitude, Longitude float64) (*models.Address, error) {
	location := &models.Location{
		Longitude: Longitude,
		Latitude:  Latitude,
	}

	address := &models.Address{
		ID:        rand.UUID(),
		EaterID:   EaterID,
		Name:      name,
		Location:  location,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}

	err := s.addressRepo.CreateAddress(ctx, address)

	if err != nil {
		return nil, err
	}

	return address, nil
}

func (s *addressSvcImpl) UpdateAddress(ctx context.Context, addressID, EaterID, name string, Latitude, Longitude float64) (*models.Address, error) {
	location := &models.Location{
		Longitude: Longitude,
		Latitude:  Latitude,
	}

	address := &models.Address{
		ID:        addressID,
		EaterID:   EaterID,
		Name:      name,
		Location:  location,
		UpdatedAt: time.Now().UTC(),
	}

	err := s.addressRepo.UpdateAddress(ctx, address)
	if err != nil {
		return nil, err
	}
	return address, nil
}

func (s *addressSvcImpl) DeleteAddress(ctx context.Context, addressID string) error {
	return s.addressRepo.DeleteAddress(ctx, addressID)
}

func (s *addressSvcImpl) GetAddressById(ctx context.Context, addressID string) (*models.Address, error) {
	address, err := s.addressRepo.GetAddressById(ctx, addressID)

	if err != nil {
		return nil, err
	}

	return address, nil
}

func (s *addressSvcImpl) ListAddressByEaterId(ctx context.Context, eaterID string) ([]*models.Address, error) {
	addresses, err := s.addressRepo.ListAddressByEaterId(ctx, eaterID)

	if err != nil {
		return nil, err
	}

	return addresses, nil
}
