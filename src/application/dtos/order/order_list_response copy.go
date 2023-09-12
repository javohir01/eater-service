package dtos

import "github.com/javohir01/eater-service/src/domain/order/models"

type OrderListResponse struct {
	Order []*models.Order
}

func NewOrderListResponse(order *models.Order) *OrderListResponse {
	return &OrderListResponse{
		Order: order,
	}
}
