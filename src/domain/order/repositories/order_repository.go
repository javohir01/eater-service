package repositories

import (
	"context"

	"github.com/javohir01/eater-service/src/domain/order/models"
)

type OrderRepository interface {
	CreateOrder(ctx context.Context, order *models.Order) (*models.Order, error)
	UpdateOrder(ctx context.Context, order *models.Order) (*models.Order, error)
	UpdateOrderStatus(ctx context.Context, orderID string, status string) (*models.Order, error)
	UpdatePaymentStatus(ctx context.Context, orderID string, paymentStatus string) (*models.Order, error)
	GetOrder(ctx context.Context, orderID string) (*models.Order, error)
	GetOrdersByEaterID(ctx context.Context, eaterID string) ([]*models.Order, error)
	DeleteOrder(ctx context.Context, orderID string) error
}
