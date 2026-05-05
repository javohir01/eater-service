package address

import (
	"context"
	"time"

	"github.com/javohir01/eater-service/src/domain/address/models"
	"github.com/javohir01/eater-service/src/infrastructure/rand"
	"gorm.io/gorm"
)

const (
	TableAddress = "eater.address"
)

type AddressRepoImpl struct {
	db *gorm.DB
}

func NewAddressRepo(db *gorm.DB) *AddressRepoImpl {
	return &AddressRepoImpl{
		db: db,
	}
}

func (r *AddressRepoImpl) WithTx(ctx context.Context, f func(r *AddressRepoImpl) error) error {
	if err := r.db.Transaction(func(tx *gorm.DB) error {
		r := &AddressRepoImpl{
			db: tx,
		}
		if err := f(r); err != nil {
			return err
		}
		return nil
	}); err != nil {
		return err
	}
	return nil
}

func (r *AddressRepoImpl) CreateAddress(ctx context.Context, name string, longitude, latitude float64) error {
	address := &models.Address{
		ID:        rand.String(10),
		Name:      name,
		Location:  &models.Location{Longitude: longitude, Latitude: latitude},
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	result := r.db.WithContext(ctx).Table(TableAddress).Create(address)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *AddressRepoImpl) UpdateAddress(ctx context.Context, addressID, name string, longitude, latitude float64) error {
	updateData := map[string]interface{}{
		"name":       name,
		"location":   &models.Location{Longitude: longitude, Latitude: latitude},
		"updated_at": time.Now(),
	}
	result := r.db.WithContext(ctx).Table(TableAddress).Where("id = ?", addressID).Updates(updateData)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
