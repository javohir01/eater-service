package dtos

import (
	"github.com/javohir01/eater-service/src/domain/rating/models"
)

type SaveRestaurantRatingResponse struct {
	RestaurantRating *models.RestaurantRating
}

func NewSaveRestaurantRatingResponse(restaurantRating *models.RestaurantRating) *SaveRestaurantRatingResponse {
	return &SaveRestaurantRatingResponse{
		RestaurantRating: restaurantRating,
	}
}
