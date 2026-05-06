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

func (r *AddressRepoImpl) CreateAddress(ctx context.Context, eaterID, name string, longitude, latitude float64) error {
	address := &models.Address{
		ID:        rand.String(10),
		EaterID:   eaterID,
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

func (r *AddressRepoImpl) DeleteAddress(ctx context.Context, addressID string) error {
	result := r.db.WithContext(ctx).Table(TableAddress).Where("id = ?", addressID).Delete(&models.Address{})
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *AddressRepoImpl) GetAddress(ctx context.Context, addressID string) (*models.Address, error) {
	var address models.Address
	result := r.db.WithContext(ctx).Table(TableAddress).First(&address, "id = ?", addressID)
	if result.Error != nil {
		return nil, result.Error
	}
	return &address, nil
}

func (r *AddressRepoImpl) GetAddressesByEaterID(ctx context.Context, eaterID string) ([]*models.Address, error) {
	var addresses []*models.Address
	result := r.db.WithContext(ctx).Table(TableAddress).Where("eater_id = ?", eaterID).Find(&addresses)
	if result.Error != nil {
		return nil, result.Error
	}
	return addresses, nil
}
