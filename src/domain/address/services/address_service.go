package services

import (
	"context"

	"github.com/javohir01/eater-service/src/domain/address/models"
	"github.com/javohir01/eater-service/src/domain/address/repositories"
	"go.uber.org/zap"
)

type AddressService interface {
	CreateAddress(ctx context.Context, address *models.Address) (string, error)
	UpdateAddress(ctx context.Context, address *models.Address)

	CreateAdress(ctx context.Context, adress *models.Address) error
	UpdateAdress(ctx context.Context, adress *models.Address) error
	DeleteAddress(ctx context.Context, id string) (*models.Address, error)
	GetAddressById(ctx context.Context, id string) (*models.Address, error)
	ListAddressByEaterId(ctx context.Context, eaterID string) (*models.Address, error)
}

type addressSvcImpl struct {
	addressRepo repositories.AddressRepository;
	logger      *zap.Logger;
}

func NewAddressService(
	addressRepo repositories.AddressRepository,
	logger *zap.Logger,
) {
	
}
