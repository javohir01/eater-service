package rating

type CreateRestaurantRatingRequest struct {
	EaterID      string    `json:"eater_id"`
	RestaurantID string    `json:"restaurant_id"`
	Rating       int       `json:"rating"`
	Comment      string    `json:"comment"` 
}
