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
