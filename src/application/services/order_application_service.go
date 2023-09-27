package services

import (
	"context"
	"errors"

	dtos "github.com/javohir01/eater-service/src/application/dtos/order"
	pb "github.com/javohir01/eater-service/src/application/protos/eater"
	addressSvc "github.com/javohir01/eater-service/src/domain/address/services"
	models "github.com/javohir01/eater-service/src/domain/order/models"
	orderSvc "github.com/javohir01/eater-service/src/domain/order/services"
)

type OrderApplicationService interface {
	CreateOrder(ctx context.Context, req *pb.PlaceOrderRequest) (*pb.PlaceOrderResponse, error)
	UpdateOrder(ctx context.Context, req *pb.UpdateOrderRequest) (*pb.UpdateOrderResponse, error)
	DeleteOrder(ctx context.Context, req *pb.DeleteOrderRequest) (*pb.DeleteOrderResponse, error)
	GetOrderById(ctx context.Context, req *pb.GetOrderRequest) (*pb.GetOrderResponse, error)
	ListOrderByEaterId(ctx context.Context, req *pb.ListOrderByEaterRequest) (*pb.ListOrderByEaterResponse, error)
}

type orderAppSvcImpl struct {
	orderSvc   orderSvc.OrderService
	addressSvc addressSvc.AddressService
}

func NewOrderApplicationService(
	orderSvc orderSvc.OrderService,
	addressSvc addressSvc.AddressService,
) OrderApplicationService {
	return &orderAppSvcImpl{
		orderSvc:   orderSvc,
		addressSvc: addressSvc,
	}
}

func (s *orderAppSvcImpl) CreateOrder(ctx context.Context, req *pb.PlaceOrderRequest) (*pb.PlaceOrderResponse, error) {
	if err := dtos.ValidateOrderRequestDB(req); err != nil {
		return nil, err
	}

	address, err := s.addressSvc.GetAddressById(ctx, req.Cart.Delivery.AddressId)
	if err != nil {
		return nil, err
	}

	addressInfo := models.AddressInfo{
		ID:        address.ID,
		Name:      address.Name,
		Longitude: address.Location.Longitude,
		Latitude:  address.Location.Latitude,
	}

	cart := dtos.ToCart(req.GetCart())

	order, err := s.orderSvc.CreateOrder(ctx, req.GetEaterId(), cart, &addressInfo)
	if err != nil {
		return nil, err
	}

	return &pb.PlaceOrderResponse{
		Order: dtos.ToOrderPB(order),
	}, nil
}

func (s *orderAppSvcImpl) UpdateOrder(ctx context.Context, req *pb.UpdateOrderRequest) (*pb.UpdateOrderResponse, error) {
	order := dtos.ToOrder(req.GetOrder())
	if err := s.orderSvc.UpdateOrder(ctx, order); err != nil {
		return nil, err
	}

	return &pb.PlaceOrderResponse{
		Order: ToOrderDb(result),
	}, nil
}

func (s *orderAppSvcImpl) DeleteOrder(ctx context.Context, orderID string) error {
	if req.GetOrderId() == "" {
		return nil, errors.New("orderID is is required!")
	}

	err := s.orderSvc.DeleteOrder(ctx, req.GetOrderId())
	if err != nil {
		return nil, err
	}

	return &pb.DeleteAddressResponse{}, nil
}

func (s *orderAppSvcImpl) GetOrderById(ctx context.Context, req *pb.GetOrderRequest) (*pb.GetOrderResponse, error) {
	if req.GetOrderId() == "" {
		return nil, errors.New("orderID is is required!")
	}

	order, err := s.orderSvc.GetOrderById(ctx, req.GetOrderId())
	if err != nil {
		return nil, err
	}

	return &pb.GetOrderResponse{
		Order: dtos.ToOrderDb(order),
	}, nil
}

func (s *orderAppSvcImpl) ListOrderByEaterId(ctx context.Context, req *pb.ListOrderByEaterRequest) (*pb.ListOrderByEaterResponse, error) {
	if req.GetOrderId() == "" {
		return nil, errors.New("orderID is is required!")
	}

	if req.GetSort() == "" {
		req.Sort = models.SortByCreatedAtDesc //default sort
	}

	if req.GetPageSize() == 0 {
		req.PageSize = 15 //default page size
	}

	orders, err := s.orderSvc.ListOrderByEaterId(ctx, req.GetOrderId(), req.GetSort(), int(req.GetPage()), int(req.GetPageSize()))
	if err != nil {
		return nil, err
	}

	return &pb.ListAddressByEaterIdResponse{
		Orders: dtos.ToOrdersDb(orders),
	}, nil
}
