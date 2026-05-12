package rating

type CreateDeliveryRatingRequest struct {
	EaterID      string    `json:"eater_id"`
	OrderID      string    `json:"order_id"`
	Rating       int       `json:"rating"`
	Comment      string    `json:"comment"`
}