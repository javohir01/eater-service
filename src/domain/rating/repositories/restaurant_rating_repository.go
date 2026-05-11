package repositories

import (
	"context"
	"github.com/javohir01/eater-service/src/domain/rating/models"
)

type RestaurantRatingRepository interface {
	CreateRestaurantRating(ctx context.Context, restaurantRating *models.RestaurantRating) (*models.RestaurantRating, error)
	UpdateRestaurantRating(ctx context.Context, restaurantRating *models.RestaurantRating) (*models.RestaurantRating, error)
	GetRestaurantRating(ctx context.Context, id string) (*models.RestaurantRating, error)
	GetRestaurantRatingsByEaterID(ctx context.Context, eaterID string) ([]*models.RestaurantRating, error)
	GetRestaurantRatingsByRestaurantID(ctx context.Context, restaurantID string) ([]*models.RestaurantRating, error)
	DeleteRestaurantRating(ctx context.Context, id string) error
}