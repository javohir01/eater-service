package order

import "github.com/javohir01/eater-service/src/domain/order/models"

type CreateOrderRequest struct {
	EaterID         string            `json:"eater_id" binding:"required"`
	Instruction     string            `json:"instruction"`
	RestaurantID    string            `json:"restaurant_id" binding:"required"`
	DeliveryInfo    DeliveryInfoDTO   `json:"delivery_info" binding:"required"`
	PaymentInfo     PaymentInfoDTO    `json:"payment_info" binding:"required"`
	Items           []*OrderItemDTO   `json:"items" binding:"required"`
}

type DeliveryInfoDTO struct {
	DeliveryAddress string `json:"delivery_address" binding:"required"`
	DeliveryTime    string `json:"delivery_time" binding:"required"`
}

type OrderItemDTO struct {
	ProductID int`json:"product_id" binding:"required"`
	Quantity  int64 `json:"quantity" binding:"required"`
	Price     int64 `json:"price" binding:"required"`
}

type PaymentInfoDTO struct {
	PaymentMethod string `json:"payment_method" binding:"required"`	
	CardID        string `json:"card_id" binding:"required"`
	DeliveryPrice int64  `json:"delivery_price" binding:"required"`
}