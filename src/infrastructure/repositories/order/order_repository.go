package order

import (
	"context"

	"github.com/javohir01/eater-service/src/domain/order/models"
	"github.com/javohir01/eater-service/src/domain/order/repositories"
	"github.com/javohir01/eater-service/src/infrastructure/utils"
	"gorm.io/gorm"
)

const (
	tableOrder      = "eater.order"
	tableCart       = "eater.cart"
	tableOrderItems = "eater.order_items"
)

type orderRepoImpl struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) repositories.OrderRepository {
	return &orderRepoImpl{
		db: db,
	}
}

func (r *orderRepoImpl) WithTx(ctx context.Context, f func(r repositories.OrderRepository) error) error {
	if err := r.db.Transaction(func(tx *gorm.DB) error {
		r := orderRepoImpl{
			db: tx,
		}
		if err := f(&r); err != nil {
			return err
		}
		return nil
	}); err != nil {
		return err
	}
	return nil
}

func (r *orderRepoImpl) SaveOrder(ctx context.Context, order *models.Order) error {
	result := r.db.WithContext(ctx).Table(tableOrder).Create(order)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *orderRepoImpl) SaveOrderItems(ctx context.Context, orderItem []*models.OrderItem) error {
	result := r.db.WithContext(ctx).Table(tableOrderItems).Create(orderItem)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *orderRepoImpl) UpdateOrder(ctx context.Context, order *models.Order) error {
	result := r.db.WithContext(ctx).Table(tableOrder).Save(order)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *orderRepoImpl) UpdateOrderStatus(ctx context.Context, orderID, status string) error {
	result := r.db.WithContext(ctx).Table(tableOrder).Where("id =?", orderID).Update("status", status)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *orderRepoImpl) UpdateOrderPaymentStatus(ctx context.Context, orderID, status string) error {
	return r.db.WithContext(ctx).Table(tableOrder).Where("id =?", orderID).Update("payment_status", status).Error
}

func (r *orderRepoImpl) GetOrder(ctx context.Context, order_id string) (*models.Order, error) {
	var order models.Order
	result := r.db.WithContext(ctx).Table(tableOrder).First(&order, "id = ?", order_id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &order, nil
}

func (r *orderRepoImpl) ListOrderByEaterId(ctx context.Context, eater_id, sort string, page, perPage int) ([]*models.Order, error) {
	var orders []*models.Order
	result := r.db.WithContext(ctx).
		Table(tableOrder).
		Where("eater_id = ?", eater_id).
		Order(sort).
		Scopes(utils.Paginate(page, perPage)).
		Find(orders)
	if result.Error != nil {
		return nil, result.Error
	}
	return &orders, nil
}

func (r *orderRepoImpl) DeleteOrder(ctx context.Context, id string) error {
	result := r.db.WithContext(ctx).Table(tableOrder).Delete("id = ?", id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
