package services

import (
	"context"
	"errors"

	dtos "github.com/javohir01/eater-service/src/application/dtos/address"
	pb "github.com/javohir01/eater-service/src/application/protos/eater"
	addressSvc "github.com/javohir01/eater-service/src/domain/address/services"
)

type AddressApplicationService interface {
	CreateAddress(ctx context.Context, req *pb.AddAddressRequest) (*pb.AddAddressResponse, error)
	UpdateAddress(ctx context.Context, req *pb.UpdateAddressRequest) (*pb.UpdateAddressRequest, error)
	DeleteAddress(ctx context.Context, req *pb.DeleteAddressRequest) (pb.UpdateAddressResponse, error)
	GetAddressById(ctx context.Context, req *pb.GetAddressRequest) (*pb.GetAddressResponse, error)
	ListAddressByEaterId(ctx context.Context, req *pb.ListAddressByEaterRequest) (*pb.ListAddressByEaterResponse, error)
}

type addressAppSvcImpl struct {
	addressSvc addressSvc.AddressService
}

func NewAddressApplicationService(
	addressSvc addressSvc.AddressService,
) AddressApplicationService {
	return &addressAppSvcImpl{
		addressSvc: addressSvc,
	}
}

func (s *addressAppSvcImpl) CreateAddress(ctx context.Context, req *pb.AddAddressRequest) (*pb.AddAddressResponse, error) {
	if req.GetEaterId() == "" {
		return nil, errors.New("EaterID is is required!")
	}

	if req.GetName() == "" {
		return nil, errors.New("name is is required!")
	}

	if req.GetLongitude() == 0 || req.GetLatitude() == 0 {
		return nil, errors.New("invalid location")
	}

	address, err := s.addressSvc.CreateAddress(ctx, req.GetEaterId(), req.GetName(), req.GetLongitude(), req.GetLatitude())
	if err != nil {
		return nil, err
	}

	return &pb.AddAddressResponse{
		Address: dtos.ToAdressPB(address),
	}, nil
}

func (s *addressAppSvcImpl) UpdateAddress(ctx context.Context, req *pb.UpdateAddressRequest) (*pb.UpdateAddressRequest, error) {
	if err := dtos.ValidateUpdateAddressRequestDB(req); err != nil {
		return nil, err
	}

	address, err := s.UpdateAddress(ctx, req.GetEaterId(), req.GetName(), req.GetLongitude(), req.GetLatitude())
	if err != nil {
		return nil, err
	}

	return &pb.UpdateAddressResponse{
		Address: dtos.ToAdressPB(address),
	}, nil
}

func (s *addressAppSvcImpl) DeleteAddress(ctx context.Context, req *pb.DeleteAddressRequest) (pb.UpdateAddressResponse, error) {
	if req.GetAddressId() == "" {
		return errors.New("addressID is is required!")
	}

	if err := s.DeleteAddress(ctx, req.GetAddressId()); err != nil {
		return err
	}

	return &pb.DeleteAddressResponse{}, nil
}

func (s *addressAppSvcImpl) GetAddressById(ctx context.Context, req *pb.GetAddressRequest) (*pb.GetAddressResponse, error) {
	if req.GetAddressId() == "" {
		return errors.New("addressID is is required!")
	}

	address, err := s.GetAddressById(ctx, req.GetAddressId())
	if err != nil {
		return nil, err
	}

	return &pb.GetAddressByIdResponse{
		Address: dtos.ToAdressPB(address),
	}, nil
}

func (s *addressAppSvcImpl) ListAddressByEaterId(ctx context.Context, req *pb.ListAddressByEaterRequest) (*pb.ListAddressByEaterResponse, error) {
	if req.GetAddressId() == "" {
		return errors.New("addressID is is required!")
	}

	addresses, err := s.ListAddressByEaterId(ctx, req.GetAddressId())
	if err != nil {
		return nil, err
	}

	return &pb.ListAddressByEaterResponse{
		Addresses: dtos.ToAdressesPB(addresses),
	}, nil
}
