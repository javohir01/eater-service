package repositories

import (
	"context"

	"github.com/javohir01/eater-service/src/domain/address/models"
)

type AddressRepository interface {
	CreateAdress(ctx context.Context, adress *models.Address) error
	UpdateAdress(ctx context.Context, adress *models.Address) error
	DeleteAddress(ctx context.Context, id string) (*models.Address, error)
	GetAddressById(ctx context.Context, id string) (*models.Address, error)
	ListAddressByEaterId(ctx context.Context, eaterID string) (*models.Address, error)
}
