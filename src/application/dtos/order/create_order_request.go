package order

import "github.com/javohir01/eater-service/src/domain/order/models"

type CreateOrderRequest struct {
	EaterID         string                    `json:"eater_id" binding:"required"`
	Instruction     string                    `json:"instruction"`
	RestaurantID    string                    `json:"restaurant_id" binding:"required"`
	DeliveryInfo    *models.DeliveryInfo      `json:"delivery_info" binding:"required"`
	PaymentInfo     *models.PaymentInfo       `json:"payment_info" binding:"required"`
	Items           []*models.OrderItem `json:"items" binding:"required"`
	Status          string                    `json:"status"`
	PaymentStatus   string                    `json:"payment_status"`
}
