package rating

import (
	"context"
	"time"

	"github.com/javohir01/eater-service/src/domain/rating/models"
	"github.com/javohir01/eater-service/src/infrastructure/rand"
	"gorm.io/gorm"
)

const (
	TableRestaurantRating = "restaurant_ratings"
)

type RestaurantRatingRepositoryImpl struct {
	db *gorm.DB
}

func NewRestaurantRatingRepository(db *gorm.DB) *RestaurantRatingRepositoryImpl {
	return &RestaurantRatingRepositoryImpl{db: db}
}

func (r *RestaurantRatingRepositoryImpl) CreateRestaurantRating(ctx context.Context, restaurantRating *models.RestaurantRating) (*models.RestaurantRating, error) {
	restaurantRating.ID = rand.String(10)
	restaurantRating.CreatedAt = time.Now()
	restaurantRating.UpdatedAt = time.Now()
	result := r.db.WithContext(ctx).Table(TableRestaurantRating).Create(restaurantRating)
	if result.Error != nil {
		return nil, result.Error
	}
	return restaurantRating, nil
}

func (r *RestaurantRatingRepositoryImpl) UpdateRestaurantRating(ctx context.Context, restaurantRating *models.RestaurantRating) (*models.RestaurantRating, error) {
	result := r.db.WithContext(ctx).Table(TableRestaurantRating).Where("id = ?", restaurantRating.ID).Updates(restaurantRating)
	if result.Error != nil {
		return nil, result.Error
	}
	return restaurantRating, nil
}

func (r *RestaurantRatingRepositoryImpl) GetRestaurantRating(ctx context.Context, id string) (*models.RestaurantRating, error) {
	var restaurantRating models.RestaurantRating
	result := r.db.WithContext(ctx).Table(TableRestaurantRating).Where("id = ?", id).First(&restaurantRating)
	if result.Error != nil {
		return nil, result.Error
	}
	return &restaurantRating, nil
}

func (r *RestaurantRatingRepositoryImpl) GetRestaurantRatingsByEaterID(ctx context.Context, eaterID string) ([]*models.RestaurantRating, error) {
	var restaurantRatings []*models.RestaurantRating
	result := r.db.WithContext(ctx).Table(TableRestaurantRating).Where("eater_id = ?", eaterID).Find(&restaurantRatings)
	if result.Error != nil {
		return nil, result.Error
	}
	return restaurantRatings, nil
}

func (r *RestaurantRatingRepositoryImpl) GetRestaurantRatingsByRestaurantID(ctx context.Context, restaurantID string) ([]*models.RestaurantRating, error) {
	var restaurantRatings []*models.RestaurantRating
	result := r.db.WithContext(ctx).Table(TableRestaurantRating).Where("restaurant_id = ?", restaurantID).Find(&restaurantRatings)
	if result.Error != nil {
		return nil, result.Error
	}
	return restaurantRatings, nil
}

func (r *RestaurantRatingRepositoryImpl) DeleteRestaurantRating(ctx context.Context, id string) error {
	result := r.db.WithContext(ctx).Table(TableRestaurantRating).Where("id = ?", id).Delete(&models.RestaurantRating{})
	return result.Error
}
