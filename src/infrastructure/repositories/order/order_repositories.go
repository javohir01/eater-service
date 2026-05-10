package order

import (
	"context"

	"github.com/javohir01/eater-service/src/domain/order/models"
	"github.com/javohir01/eater-service/src/domain/order/repositories"
	"github.com/javohir01/eater-service/src/utils/rand"
	"gorm.io/gorm"
)

const TableOrder = "orders"

type OrderRepositoryImpl struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) *OrderRepositoryImpl {
	return &OrderRepositoryImpl{db: db}
}

func WithTx(ctx context.Context, db *gorm.DB, f func(r repositories.OrderRepository) error) error {
	return db.Transaction(func(tx *gorm.DB) error {
		r := &OrderRepositoryImpl{db: tx}
		return f(r)
	})
}

func (r *OrderRepositoryImpl) CreateOrder(ctx context.Context, order *models.Order) (*models.Order, error) {
	id := rand.String(10)
	order.ID = id
	result := r.db.WithContext(ctx).Table(TableOrder).Create(order)
	if result.Error != nil {
		return nil, result.Error
	}
	return order, nil
}

func (r *OrderRepositoryImpl) UpdateOrder(ctx context.Context, order *models.Order) (*models.Order, error) {
	result := r.db.WithContext(ctx).Table(TableOrder).Where("id = ?", order.ID).Updates(order)
	if result.Error != nil {
		return nil, result.Error
	}
	return order, nil
}

func (r *OrderRepositoryImpl) UpdateOrderStatus(ctx context.Context, orderID string, status string) (*models.Order, error) {
	result := r.db.WithContext(ctx).Table(TableOrder).Where("id = ?", orderID).Update("status", status)
	if result.Error != nil {
		return nil, result.Error
	}
	return r.GetOrder(ctx, orderID)
}

func (r *OrderRepositoryImpl) UpdatePaymentStatus(ctx context.Context, orderID string, paymentStatus string) (*models.Order, error) {
	result := r.db.WithContext(ctx).Table(TableOrder).Where("id = ?", orderID).Update("payment_status", paymentStatus)
	if result.Error != nil {
		return nil, result.Error
	}
	return r.GetOrderByID(ctx, orderID)
}

func (r *OrderRepositoryImpl) GetOrder(ctx context.Context, id string) (*models.Order, error) {
	var order models.Order
	result := r.db.WithContext(ctx).Table(TableOrder).Where("id = ?", id).First(&order)
	if result.Error != nil {
		return nil, result.Error
	}
	return &order, nil
}

func (r *OrderRepositoryImpl) GetOrdersByUserID(ctx context.Context, userID string) ([]*models.Order, error) {
	var orders []*models.Order
	result := r.db.WithContext(ctx).Table(TableOrder).Where("user_id = ?", userID).Find(&orders)
	if result.Error != nil {
		return nil, result.Error
	}
	return orders, nil
}

func (r *OrderRepositoryImpl) DeleteOrder(ctx context.Context, id string) error {
	result := r.db.WithContext(ctx).Table(TableOrder).Where("id = ?", id).Delete(&models.Order{})
	return result.Error
}
