package services

import (
	"context"

	"github.com/javohir01/eater-service/src/application/dtos/rating"
	dtos "github.com/javohir01/eater-service/src/application/dtos/rating"
	"github.com/javohir01/eater-service/src/domain/rating/models"
	"github.com/javohir01/eater-service/src/domain/rating/services"
)

type RatingService interface {
	CreateRestaurantRating(ctx context.Context, dto *dtos.CreateRestaurantRatingRequest) (*dtos.RestaurantRatingResponse, error)
	UpdateRestaurantRating(ctx context.Context, dto *dtos.UpdateRestaurantRatingRequest) (*dtos.RestaurantRatingResponse, error)
	GetRestaurantRating(ctx context.Context, id string) (*dtos.RestaurantRatingResponse, error)
	GetRestaurantRatingsByEaterID(ctx context.Context, eaterID string) ([]*dtos.RestaurantRatingResponse, error)
	GetRestaurantRatingsByRestaurantID(ctx context.Context, restaurantID string) ([]*dtos.RestaurantRatingResponse, error)
	DeleteRestaurantRating(ctx context.Context, id string) error

	CreateDeliveryRating(ctx context.Context, rating *models.DeliveryRating) (*models.DeliveryRating, error)
	UpdateDeliveryRating(ctx context.Context, rating *models.DeliveryRating) (*models.DeliveryRating, error)
	GetDeliveryRating(ctx context.Context, id string) (*models.DeliveryRating, error)
	GetDeliveryRatingsByEaterID(ctx context.Context, eaterID string) ([]*models.DeliveryRating, error)
	GetDeliveryRatingsByOrderID(ctx context.Context, orderID string) ([]*models.DeliveryRating, error)
	DeleteDeliveryRating(ctx context.Context, id string) error
}

type RatingServiceImpl struct {
	ratingService services.RatingService
}

func NewRatingService(ratingService services.RatingService) *RatingServiceImpl {
	return &RatingServiceImpl{ratingService: ratingService}
}

func (s *RatingServiceImpl) CreateRestaurantRating(ctx context.Context, dto *dtos.CreateRestaurantRatingRequest) (*dtos.RestaurantRatingResponse, error) {
	model := models.RestaurantRating{
		EaterID:      dto.EaterID,
		RestaurantID: dto.RestaurantID,
		Rating:       dto.Rating,
		Comment:      dto.Comment,
	}
	result, err := s.ratingService.CreateRestaurantRating(ctx, &model)
	if err != nil {
		return nil, err
	}
	return &dtos.RestaurantRatingResponse{
		ID:           result.ID,
		EaterID:      result.EaterID,
		RestaurantID: result.RestaurantID,
		Rating:       result.Rating,
		Comment:      result.Comment,
	}, nil
}

func (s *RatingServiceImpl) UpdateRestaurantRating(ctx context.Context, dto *dtos.UpdateRestaurantRatingRequest) (*dtos.RestaurantRatingResponse, error) {
	model := models.RestaurantRating{
		ID:           dto.ID,
		EaterID:      dto.EaterID,
		RestaurantID: dto.RestaurantID,
		Rating:       dto.Rating,
		Comment:      dto.Comment,
	}
	result, err := s.ratingService.UpdateRestaurantRating(ctx, &model)
	if err != nil {
		return nil, err
	}
	return &dtos.RestaurantRatingResponse{
		ID:           result.ID,
		EaterID:      result.EaterID,
		RestaurantID: result.RestaurantID,
		Rating:       result.Rating,
		Comment:      result.Comment,
	}, nil
}

func (s *RatingServiceImpl) GetRestaurantRating(ctx context.Context, id string) (*dtos.RestaurantRatingResponse, error) {
	result, err := s.ratingService.GetRestaurantRating(ctx, id)
	if err != nil {
		return nil, err
	}
	return &dtos.RestaurantRatingResponse{
		ID:           result.ID,
		EaterID:      result.EaterID,
		RestaurantID: result.RestaurantID,
		Rating:       result.Rating,
		Comment:      result.Comment,
	}, nil
}

func (s *RatingServiceImpl) GetRestaurantRatingsByEaterID(ctx context.Context, eaterID string) ([]*dtos.RestaurantRatingResponse, error) {
	result, err := s.ratingService.GetRestaurantRatingsByEaterID(ctx, eaterID)
	if err != nil {
		return nil, err
	}
	var restaurantRatings []*dtos.RestaurantRatingResponse
	for _, rating := range result {
		ratingResponse := dtos.RestaurantRatingResponse{
			ID:           rating.ID,
			EaterID:      rating.EaterID,
			RestaurantID: rating.RestaurantID,
			Rating:       rating.Rating,
			Comment:      rating.Comment,
		}
		restaurantRatings = append(restaurantRatings, &ratingResponse)
	}
	return restaurantRatings, nil
}

func (s *RatingServiceImpl) GetRestaurantRatingsByRestaurantID(ctx context.Context, restaurantID string) ([]*dtos.RestaurantRatingResponse, error) {
	result, err := s.ratingService.GetRestaurantRatingsByRestaurantID(ctx, restaurantID)
	if err != nil {
		return nil, err
	}
	var restaurantRatings []*dtos.RestaurantRatingResponse
	for _, rating := range result {
		ratingResponse := dtos.RestaurantRatingResponse{
			ID:           rating.ID,
			EaterID:      rating.EaterID,
			RestaurantID: rating.RestaurantID,
			Rating:       rating.Rating,
			Comment:      rating.Comment,
		}
		restaurantRatings = append(restaurantRatings, &ratingResponse)
	}
	return restaurantRatings, nil
}

func (s *RatingServiceImpl) DeleteRestaurantRating(ctx context.Context, id string) error {
	return s.ratingService.DeleteRestaurantRating(ctx, id)
}

func (s *RatingServiceImpl) CreateDeliveryRating(ctx context.Context, dto *dtos.CreateDeliveryRatingRequest) (*dtos.DeliveryRatingResponse, error) {
	model := models.DeliveryRating{
		EaterID: dto.EaterID,
		OrderID: dto.OrderID,
		Rating:  dto.Rating,
		Comment: dto.Comment,
	}
	result, err := s.ratingService.CreateDeliveryRating(ctx, &model)
	if err != nil {
		return nil, err
	}
	return &dtos.DeliveryRatingResponse{
		ID:      result.ID,
		EaterID: result.EaterID,
		OrderID: result.OrderID,
		Rating:  result.Rating,
		Comment: result.Comment,
	}, nil
}

func (s *RatingServiceImpl) UpdateDeliveryRating(ctx context.Context, dto *rating.UpdateDeliveryRatingRequest) (*rating.DeliveryRatingResponse, error) {
	model := models.DeliveryRating{
		ID:      dto.ID,
		EaterID: dto.EaterID,
		OrderID: dto.OrderID,
		Rating:  dto.Rating,
		Comment: dto.Comment,
	}
	result, err := s.ratingService.UpdateDeliveryRating(ctx, &model)
	if err != nil {
		return nil, err
	}
	return &dtos.DeliveryRatingResponse{
		ID:      result.ID,
		EaterID: result.EaterID,
		OrderID: result.OrderID,
		Rating:  result.Rating,
		Comment: result.Comment,
	}, nil
}

func (s *RatingServiceImpl) GetDeliveryRating(ctx context.Context, id string) (*dtos.DeliveryRatingResponse, error) {
	result, err := s.ratingService.GetDeliveryRating(ctx, id)
	if err != nil {
		return nil, err
	}
	return &dtos.DeliveryRatingResponse{
		ID:      result.ID,
		EaterID: result.EaterID,
		OrderID: result.OrderID,
		Rating:  result.Rating,
		Comment: result.Comment,
	}, nil
}

func (s *RatingServiceImpl) GetDeliveryRatingsByEaterID(ctx context.Context, eaterID string) ([]*dtos.DeliveryRatingResponse, error) {
	result, err := s.ratingService.GetDeliveryRatingsByEaterID(ctx, eaterID)
	if err != nil {
		return nil, err
	}
	var deliveryRatings []*dtos.DeliveryRatingResponse
	for _, rating := range result {
		deliveryRatingResponse := dtos.DeliveryRatingResponse{
			ID:      rating.ID,
			EaterID: rating.EaterID,
			OrderID: rating.OrderID,
			Rating:  rating.Rating,
			Comment: rating.Comment,
		}
		deliveryRatings = append(deliveryRatings, &deliveryRatingResponse)
	}
	return deliveryRatings, nil
}

func (s *RatingServiceImpl) GetDeliveryRatingsByOrderID(ctx context.Context, orderID string) ([]*rating.DeliveryRatingResponse, error) {
	result, err := s.ratingService.GetDeliveryRatingsByOrderID(ctx, orderID)
	if err != nil {
		return nil, err
	}
	var deliveryRatings []*dtos.DeliveryRatingResponse
	for _, rating := range result {
		deliveryRatingResponse := dtos.DeliveryRatingResponse{
			ID:      rating.ID,
			EaterID: rating.EaterID,
			OrderID: rating.OrderID,
			Rating:  rating.Rating,
			Comment: rating.Comment,
		}
		deliveryRatings = append(deliveryRatings, &deliveryRatingResponse)
	}
	return deliveryRatings, nil
}

func (s *RatingServiceImpl) DeleteDeliveryRating(ctx context.Context, id string) error {
	return s.ratingService.DeleteDeliveryRating(ctx, id)
}
