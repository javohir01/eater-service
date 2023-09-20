package repositories

import (
	"context"

	"github.com/javohir01/eater-service/src/domain/order/models"
)

type OrderRepository interface {
	WithTx(ctx context.Context, f func(r OrderRepository) error) error
	SaveOrder(ctx context.Context, order *models.Order) error
	SaveOrderItems(ctx context.Context, orderItem []*models.OrderItem) error
	UpdateOrder(ctx context.Context, order *models.Order) error
	UpdateOrderStatus(ctx context.Context, orderID, status string) error
	UpdateOrderPaymentStatus(ctx context.Context, orderID, status string) error
	GetOrder(ctx context.Context, orderID string) (*models.Order, error)
	ListOrderByEaterId(ctx context.Context, eater_id, sort string, page, perPage int) ([]*models.Order, error)
	DeleteAddress(ctx context.Context, id string) error
}
