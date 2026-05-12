package rating

type RestaurantRatingResponse struct {
	ID           string    `json:"id"`
	EaterID      string    `json:"eater_id"`
	RestaurantID string    `json:"restaurant_id"`
	Rating       int       `json:"rating"`
	Comment      string    `json:"comment"` 
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}