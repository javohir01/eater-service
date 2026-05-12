package rating

import (
	"context"
	"time"

	"github.com/javohir01/eater-service/src/domain/rating/models"
	"github.com/javohir01/eater-service/src/infrastructure/rand"
	"gorm.io/gorm"
)

const (
	TableDeliveryRating = "delivery_ratings"
)

type DeliveryRatingRepositoryImpl struct {
	db *gorm.DB
}

func NewDeliveryRatingRepository(db *gorm.DB) *DeliveryRatingRepositoryImpl {
	return &DeliveryRatingRepositoryImpl{db: db}
}

func (r *DeliveryRatingRepositoryImpl) CreateDeliveryRating(ctx context.Context, deliveryRating *models.DeliveryRating) (*models.DeliveryRating, error) {
	deliveryRating.ID = rand.String(10)
	deliveryRating.CreatedAt = time.Now()
	deliveryRating.UpdatedAt = time.Now()
	result := r.db.WithContext(ctx).Table(TableDeliveryRating).Create(deliveryRating)
	if result.Error != nil {
		return nil, result.Error
	}
	return deliveryRating, nil
}

func (r *DeliveryRatingRepositoryImpl) UpdateDeliveryRating(ctx context.Context, deliveryRating *models.DeliveryRating) (*models.DeliveryRating, error) {
	result := r.db.WithContext(ctx).Table(TableDeliveryRating).Where("id = ?", deliveryRating.ID).Updates(deliveryRating)
	if result.Error != nil {
		return nil, result.Error
	}
	return deliveryRating, nil
}

func (r *DeliveryRatingRepositoryImpl) GetDeliveryRating(ctx context.Context, id string) (*models.DeliveryRating, error) {
	var deliveryRating models.DeliveryRating
	result := r.db.WithContext(ctx).Table(TableDeliveryRating).Where("id = ?", id).First(&deliveryRating)
	if result.Error != nil {
		return nil, result.Error
	}
	return &deliveryRating, nil
}

func (r *DeliveryRatingRepositoryImpl) GetDeliveryRatingsByEaterID(ctx context.Context, eaterID string) ([]*models.DeliveryRating, error) {
	var deliveryRatings []*models.DeliveryRating
	result := r.db.WithContext(ctx).Table(TableDeliveryRating).Where("eater_id = ?", eaterID).Find(&deliveryRatings)
	if result.Error != nil {
		return nil, result.Error
	}
	return deliveryRatings, nil
}

func (r *DeliveryRatingRepositoryImpl) GetDeliveryRatingsByOrderID(ctx context.Context, orderID string) ([]*models.DeliveryRating, error) {
	var deliveryRatings []*models.DeliveryRating
	result := r.db.WithContext(ctx).Table(TableDeliveryRating).Where("order_id = ?", orderID).Find(&deliveryRatings)
	if result.Error != nil {
		return nil, result.Error
	}
	return deliveryRatings, nil
}

func (r *DeliveryRatingRepositoryImpl) DeleteDeliveryRating(ctx context.Context, id string) error {
	result := r.db.WithContext(ctx).Table(TableDeliveryRating).Where("id = ?", id).Delete(&models.DeliveryRating{})
	return result.Error
}
