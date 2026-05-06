package repositories

import (
	"context"

	"github.com/javohir01/eater-service/src/domain/address/models"
)

type AddressRepository interface {
	CreateAddress(ctx context.Context, eaterID, name string, longitude, latitude float64) error
	UpdateAddress(ctx context.Context, addressID, name string, longitude, latitude float64) error
	DeleteAddress(ctx context.Context, addressID string) error
	GetAddress(ctx context.Context, addressID string) (*models.Address, error)
	GetAddressesByEaterID(ctx context.Context, eaterID string) ([]*models.Address, error)
}
