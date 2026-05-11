package services

import (
	"context"

	"github.com/javohir01/eater-service/src/domain/order/models"
	"github.com/javohir01/eater-service/src/domain/order/repositories"
)

type OrderServiceImpl struct {
	repo repositories.OrderRepository
}

func NewOrderService(repo repositories.OrderRepository) *OrderServiceImpl {
	return &OrderServiceImpl{repo: repo}
}

func (s *OrderServiceImpl) CreateOrder(ctx context.Context, order *models.Order) (*models.Order, error) {
	return s.repo.CreateOrder(ctx, order)
}

func (s *OrderServiceImpl) UpdateOrder(ctx context.Context, order *models.Order) (*models.Order, error) {
	return s.repo.UpdateOrder(ctx, order)
}

func (s *OrderServiceImpl) UpdateOrderStatus(ctx context.Context, orderID string, status string) (*models.Order, error) {
	return s.repo.UpdateOrderStatus(ctx, orderID, status)
}

func (s *OrderServiceImpl) UpdatePaymentStatus(ctx context.Context, orderID string, paymentStatus string) (*models.Order, error) {
	return s.repo.UpdatePaymentStatus(ctx, orderID, paymentStatus)
}

func (s *OrderServiceImpl) GetOrder(ctx context.Context, id string) (*models.Order, error) {
	return s.repo.GetOrder(ctx, id)
}

func (s *OrderServiceImpl) GetOrdersByEaterID(ctx context.Context, eaterID string) ([]*models.Order, error) {
	return s.repo.GetOrdersByEaterID(ctx, eaterID)
}

func (s *OrderServiceImpl) DeleteOrder(ctx context.Context, id string) error {
	return s.repo.DeleteOrder(ctx, id)
}