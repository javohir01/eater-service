package dtos

import (
	"errors"

	pb "github.com/javohir01/eater-service/src/application/protos/eater"
	"github.com/javohir01/eater-service/src/domain/order/models"
)

func ValidateOrderRequestDB(req *pb.PlaceOrderRequest) error {
	if req.GetEaterId() == "" {
		return errors.New("EaterId is is required!")
	}
	if req.GetCart() == nil {
		return errors.New("cart is is required!")
	}
	if req.GetCart().GetRestaurant() == nil {
		return errors.New("cart restaurant is is required!")
	}
	if req.GetCart().GetRestaurant().GetEntityId() == "" || req.GetCart().GetRestaurant().GetRestaurantId() == "" {
		return errors.New("cart restaurant is is required!")
	}
	if req.GetCart().GetDelivery() == nil {
		return errors.New("cart restaurant is is required!")
	}
	if req.GetCart().GetDelivery().GetAddress() == nil {
		return errors.New("cart restaurant is is required!")
	}
	if req.GetCart().GetDelivery().GetAddress().GetId() == "" {
		return errors.New("cart restaurant is is required!")
	}
	return nil
}

func ToCart(cart *models.Cart) *pb.Cart {
	restaurant := &pb.Restaurant{
		EntityId:     cart.Restaurant.EntityId,
		RestaurantId: cart.Restaurant.RestaurantId,
	  }
	  delivery := &pb.Delivery{
		Address: &pb.Address{
		  Id: cart.Delivery.Address.Id,
		},
	  }
	Payment := &pb.Payment{
		Method: cart.Payment.Method,
	}
	return &pb.Cart{
		Instruction: cart.Instruction,
		Restaurant:  restaurant,
		Delivery:    delivery,
		Payment:     &pb.Payment{CardId: cart.Payment.CardID},
		Items:       make([]*pb.CartItem, 0), // Initialize an empty slice
	  }
}