package services	

import (
	"context"
	"github.com/javohir01/eater-service/src/domain/rating/models"
	"github.com/javohir01/eater-service/src/domain/rating/repositories"
)

type RatingService interface {
	CreateRestaurantRating(ctx context.Context, restaurantRating *models.RestaurantRating) (*models.RestaurantRating, error)
	UpdateRestaurantRating(ctx context.Context, restaurantRating *models.RestaurantRating) (*models.RestaurantRating, error)
	GetRestaurantRating(ctx context.Context, id string) (*models.RestaurantRating, error)
	GetRestaurantRatingsByEaterID(ctx context.Context, eaterID string) ([]*models.RestaurantRating, error)
	GetRestaurantRatingsByRestaurantID(ctx context.Context, restaurantID string) ([]*models.RestaurantRating, error)
	DeleteRestaurantRating(ctx context.Context, id string) error
	CreateDeliveryRating(ctx context.Context, deliveryRating *models.DeliveryRating) (*models.DeliveryRating, error)
	UpdateDeliveryRating(ctx context.Context, deliveryRating *models.DeliveryRating) (*models.DeliveryRating, error)
	GetDeliveryRating(ctx context.Context, id string) (*models.DeliveryRating, error)
	GetDeliveryRatingsByEaterID(ctx context.Context, eaterID string) ([]*models.DeliveryRating, error)
	GetDeliveryRatingsByOrderID(ctx context.Context, orderID string) ([]*models.DeliveryRating, error)
	DeleteDeliveryRating(ctx context.Context, id string) error
}	

type RatingServiceImpl struct {
	restaurantRepository *repositories.RestaurantRatingRepositoryImpl
	deliveryRepository *repositories.DeliveryRatingRepositoryImpl
}	

func NewRatingService(restaurantRepository *repositories.RestaurantRatingRepositoryImpl, deliveryRepository *repositories.DeliveryRatingRepositoryImpl) *RatingServiceImpl {
	return &RatingServiceImpl{restaurantRepository: restaurantRepository, deliveryRepository: deliveryRepository}
}	

func (r *RatingServiceImpl) CreateRestaurantRating(ctx context.Context, restaurantRating *models.RestaurantRating) (*models.RestaurantRating, error) {
	return r.restaurantRepository.CreateRestaurantRating(ctx, restaurantRating)
}	

func (r *RatingServiceImpl) UpdateRestaurantRating(ctx context.Context, restaurantRating *models.RestaurantRating) (*models.RestaurantRating, error) {
	return r.restaurantRepository.UpdateRestaurantRating(ctx, restaurantRating)
}	

func (r *RatingServiceImpl) GetRestaurantRating(ctx context.Context, id string) (*models.RestaurantRating, error) {
	return r.restaurantRepository.GetRestaurantRating(ctx, id)
}	

func (r *RatingServiceImpl) GetRestaurantRatingsByEaterID(ctx context.Context, eaterID string) ([]*models.RestaurantRating, error) {
	return r.restaurantRepository.GetRestaurantRatingsByEaterID(ctx, eaterID)
}	

func (r *RatingServiceImpl) GetRestaurantRatingsByRestaurantID(ctx context.Context, restaurantID string) ([]*models.RestaurantRating, error) {
	return r.restaurantRepository.GetRestaurantRatingsByRestaurantID(ctx, restaurantID)
}	

func (r *RatingServiceImpl) DeleteRestaurantRating(ctx context.Context, id string) error {
	return r.restaurantRepository.DeleteRestaurantRating(ctx, id)
}	

func (r *RatingServiceImpl) CreateDeliveryRating(ctx context.Context, deliveryRating *models.DeliveryRating) (*models.DeliveryRating, error) {
	return r.deliveryRepository.CreateDeliveryRating(ctx, deliveryRating)
}	

func (r *RatingServiceImpl) UpdateDeliveryRating(ctx context.Context, deliveryRating *models.DeliveryRating) (*models.DeliveryRating, error) {
	return r.deliveryRepository.UpdateDeliveryRating(ctx, deliveryRating)
}	

func (r *RatingServiceImpl) GetDeliveryRating(ctx context.Context, id string) (*models.DeliveryRating, error) {
	return r.deliveryRepository.GetDeliveryRating(ctx, id)
}	

func (r *RatingServiceImpl) GetDeliveryRatingsByEaterID(ctx context.Context, eaterID string) ([]*models.DeliveryRating, error) {
	return r.deliveryRepository.GetDeliveryRatingsByEaterID(ctx, eaterID)
}	

func (r *RatingServiceImpl) GetDeliveryRatingsByOrderID(ctx context.Context, orderID string) ([]*models.DeliveryRating, error) {
	return r.deliveryRepository.GetDeliveryRatingsByOrderID(ctx, orderID)
}	

func (r *RatingServiceImpl) DeleteDeliveryRating(ctx context.Context, id string) error {
	return r.deliveryRepository.DeleteDeliveryRating(ctx, id)
}