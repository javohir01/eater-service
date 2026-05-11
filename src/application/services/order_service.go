package services

import (
	"context"
	"github.com/javohir01/eater-service/src/domain/order/models"
	"github.com/javohir01/eater-service/src/domain/order/repositories"
	"github.com/javohir01/eater-service/src/infrastructure/repositories/order"
	"github.com/javohir01/eater-service/src/application/dtos/order"
)

type OrderService interface {
	CreateOrder(ctx context.Context, req *order.CreateOrderRequest) (order.OrderResponse, error)
	UpdateOrderStatus(ctx context.Context, req *order.UpdateOrderStatusRequest) (order.OrderResponse, error)
	UpdatePaymentStatus(ctx context.Context, req *order.UpdatePaymentStatusRequest) (order.OrderResponse, error)
	GetOrder(ctx context.Context, id string) (order.OrderResponse, error)
	GetOrdersByEaterID(ctx context.Context, eaterID string) ([]order.OrderResponse, error)
	DeleteOrder(ctx context.Context, id string) error
}

type OrderServiceImpl struct {
	orderRepo repositories.OrderRepository
}

func NewOrderService(orderRepo repositories.OrderRepository) OrderService {
	return &OrderServiceImpl{orderRepo: orderRepo}
}

func (s *OrderServiceImpl) CreateOrder(ctx context.Context, req *order.CreateOrderRequest) (*models.Order, error) {
	return s.orderRepo.CreateOrder(ctx, &models.Order{
		EaterID:       req.EaterID,
		Instruction:   req.Instruction,
		RestaurantID:  req.RestaurantID,
		DeliveryInfo:  req.DeliveryInfo,
		PaymentInfo:   req.PaymentInfo,
		Items:         req.Items,
		Status:        req.Status,
		PaymentStatus: req.PaymentStatus,
	})
}

func (s *OrderServiceImpl) UpdateOrderStatus(ctx context.Context, req *order.UpdateOrderStatusRequest) (*models.Order, error) {
	return s.orderRepo.UpdateOrderStatus(ctx, req.OrderID, req.Status)
}

func (s *OrderServiceImpl) UpdatePaymentStatus(ctx context.Context, req *order.UpdatePaymentStatusRequest) (*models.Order, error) {
	return s.orderRepo.UpdatePaymentStatus(ctx, req.OrderID, req.PaymentStatus)
}

func (s *OrderServiceImpl) GetOrder(ctx context.Context, id string) (*models.Order, error) {
	return s.orderRepo.GetOrder(ctx, id)
}

func (s *OrderServiceImpl) GetOrdersByEaterID(ctx context.Context, eaterID string) ([]*models.Order, error) {
	return s.orderRepo.GetOrdersByEaterID(ctx, eaterID)
}

func (s *OrderServiceImpl) DeleteOrder(ctx context.Context, id string) error {
	return s.orderRepo.DeleteOrder(ctx, id)
}	