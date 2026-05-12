package rating

type UpdateRestaurantRatingRequest struct {
	ID           string    `json:"id"`
	EaterID      string    `json:"eater_id"`
	RestaurantID string    `json:"restaurant_id"`
	Rating       int       `json:"rating"`
	Comment      string    `json:"comment"`
}