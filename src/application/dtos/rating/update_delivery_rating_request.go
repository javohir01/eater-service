package rating

type UpdateDeliveryRatingRequest struct {
	ID           string    `json:"id"`
	EaterID      string    `json:"eater_id"`
	OrderID      string    `json:"order_id"`
	Rating       int       `json:"rating"`
	Comment      string    `json:"comment"`
}