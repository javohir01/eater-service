package repositories

import (
	"context"
	"github.com/javohir01/eater-service/src/domain/rating/models"
)

type DeliveryRatingRepository interface {
	CreateDeliveryRating(ctx context.Context, deliveryRating *models.DeliveryRating) (*models.DeliveryRating, error)
	UpdateDeliveryRating(ctx context.Context, deliveryRating *models.DeliveryRating) (*models.DeliveryRating, error)
	GetDeliveryRating(ctx context.Context, id string) (*models.DeliveryRating, error)
	GetDeliveryRatingsByEaterID(ctx context.Context, eaterID string) ([]*models.DeliveryRating, error)
	GetDeliveryRatingsByOrderID(ctx context.Context, orderID string) ([]*models.DeliveryRating, error)
	DeleteDeliveryRating(ctx context.Context, id string) error
}
