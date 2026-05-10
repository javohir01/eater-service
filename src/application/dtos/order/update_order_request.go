package order

type UpdateOrderRequest struct {
	ID            string                 `json:"id" binding:"required"`
	EaterID       string                 `json:"eater_id" binding:"required"`
	Instruction   string                 `json:"instruction"`
	Restaurant    *models.RestaurantInfo `json:"restaurant"`
	Delivery      *models.DeliveryInfo   `json:"delivery"`
	Payment       *models.PaymentInfo    `json:"payment"`
	Items         []*models.OrderItem    `json:"items"`
	Status        string                 `json:"status"`
	PaymentStatus string                 `json:"payment_status"`
}