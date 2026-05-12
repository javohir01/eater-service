package services

import (
	"context"
	"strconv"
	"github.com/javohir01/eater-service/src/domain/order/models"
	"github.com/javohir01/eater-service/src/domain/order/repositories"
	dtos "github.com/javohir01/eater-service/src/application/dtos/order"
)

type OrderService interface {
	CreateOrder(ctx context.Context, req *dtos.CreateOrderRequest) (dtos.OrderResponse, error)
	UpdateOrderStatus(ctx context.Context, req *dtos.UpdateOrderStatusRequest) (dtos.OrderResponse, error)
	UpdatePaymentStatus(ctx context.Context, req *dtos.UpdatePaymentStatusRequest) (dtos.OrderResponse, error)
	GetOrder(ctx context.Context, id string) (dtos.OrderResponse, error)
	GetOrdersByEaterID(ctx context.Context, eaterID string) ([]dtos.OrderResponse, error)
	DeleteOrder(ctx context.Context, id string) error
}

type OrderServiceImpl struct {
	orderRepo repositories.OrderRepository
}

func NewOrderService(orderRepo repositories.OrderRepository) OrderService {
	return &OrderServiceImpl{orderRepo: orderRepo}
}

func (s *OrderServiceImpl) CreateOrder(ctx context.Context, req *dtos.CreateOrderRequest) (dtos.OrderResponse, error) {
	var items []*models.OrderItem
	for _, item := range req.Items {
		items = append(items, &models.OrderItem{
			ProductID: strconv.Itoa(item.ProductID),
			Quantity:  int(item.Quantity),
			Price:     item.Price,
		})
	}

	order, err := s.orderRepo.CreateOrder(ctx, &models.Order{
		EaterID:      req.EaterID,
		Instruction:  req.Instruction,
		RestaurantID: req.RestaurantID,
		Delivery: &models.DeliveryInfo{
			Address: &models.AddressInfo{Name: req.DeliveryInfo.DeliveryAddress},
			Time:    req.DeliveryInfo.DeliveryTime,
		},
		Payment: &models.PaymentInfo{
			Method:        req.PaymentInfo.PaymentMethod,
			CardID:        req.PaymentInfo.CardID,
			DeliveryPrice: req.PaymentInfo.DeliveryPrice,
		},
		Items:         items,
		Status:        "PENDING",
		PaymentStatus: "UNPAID",
	})
	if err != nil {
		return dtos.OrderResponse{}, err
	}
	return s.mapToOrderResponse(order), nil
}

func (s *OrderServiceImpl) UpdateOrderStatus(ctx context.Context, req *dtos.UpdateOrderStatusRequest) (dtos.OrderResponse, error) {
	order, err := s.orderRepo.UpdateOrderStatus(ctx, req.OrderID, req.Status)
	if err != nil {
		return dtos.OrderResponse{}, err
	}
	return s.mapToOrderResponse(order), nil
}

func (s *OrderServiceImpl) UpdatePaymentStatus(ctx context.Context, req *dtos.UpdatePaymentStatusRequest) (dtos.OrderResponse, error) {
	order, err := s.orderRepo.UpdatePaymentStatus(ctx, req.OrderID, req.PaymentStatus)
	if err != nil {
		return dtos.OrderResponse{}, err
	}
	return s.mapToOrderResponse(order), nil
}

func (s *OrderServiceImpl) GetOrder(ctx context.Context, id string) (dtos.OrderResponse, error) {
	order, err := s.orderRepo.GetOrder(ctx, id)
	if err != nil {
		return dtos.OrderResponse{}, err
	}
	return s.mapToOrderResponse(order), nil
}

func (s *OrderServiceImpl) GetOrdersByEaterID(ctx context.Context, eaterID string) ([]dtos.OrderResponse, error) {
	orders, err := s.orderRepo.GetOrdersByEaterID(ctx, eaterID)
	if err != nil {
		return nil, err
	}
	var res []dtos.OrderResponse
	for _, o := range orders {
		res = append(res, s.mapToOrderResponse(o))
	}
	return res, nil
}

func (s *OrderServiceImpl) DeleteOrder(ctx context.Context, id string) error {
	return s.orderRepo.DeleteOrder(ctx, id)
}

func (s *OrderServiceImpl) mapToOrderResponse(order *models.Order) dtos.OrderResponse {
	return dtos.OrderResponse{
		ID:            order.ID,
		EaterID:       order.EaterID,
		Instruction:   order.Instruction,
		Restaurant:    order.Restaurant,
		Delivery:      order.Delivery,
		Payment:       order.Payment,
		Items:         order.Items,
		Status:        order.Status,
		PaymentStatus: order.PaymentStatus,
		CreatedAt:     order.CreatedAt,
		UpdatedAt:     order.UpdatedAt,
	}
}