package order

type UpdatePaymentStatusRequest struct {
	OrderID       string `json:"order_id" binding:"required"`
	PaymentStatus string `json:"payment_status" binding:"required"`
}