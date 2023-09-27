package services

import (
	"context"
	"errors"
	"fmt"

	dtos "github.com/javohir01/eater-service/src/application/dtos/eater"
	pb "github.com/javohir01/eater-service/src/application/protos/eater"
	eatersvc "github.com/javohir01/eater-service/src/domain/eater/services"
	"github.com/javohir01/eater-service/src/infrastructure/jwt"
	"github.com/javohir01/eater-service/src/infrastructure/validator"
)

type EaterApplicationService interface {
	SignupEater(ctx context.Context, req *pb.SignupEaterRequest) (*pb.SignupEaterResponse, error)
	ConfirmSMSCode(ctx context.Context, req *pb.ConfirmSmsCodeRequest) (*pb.ConfirmSmsCodeResponse, error)
	UpdateEaterProfile(ctx context.Context, req *pb.UpdateEaterProfileRequest) (*pb.UpdateEaterProfileResponse, error)
	GetEaterProfile(ctx context.Context, req *pb.GetEaterProfileRequest) (*pb.GetEaterProfileResponse, error)
}

type eaterAppSvcImpl struct {
	eaterSvc eatersvc.EaterService
	tokenSvc jwt.Service
}

func NewEaterApplicationService(
	eaterSvc eatersvc.EaterService,
	tokenSvc jwt.Service,
) EaterApplicationService {
	return &eaterAppSvcImpl{
		eaterSvc: eaterSvc,
		tokenSvc: tokenSvc,
	}
}

func (s *eaterAppSvcImpl) SignupEater(ctx context.Context, req *pb.SignupEaterRequest) (*pb.SignupEaterResponse, error) {
	if !validator.ValidateUzPhoneNumber(req.GetPhoneNumber()) {
		return nil, errors.New("Invalid number. Must be off 998{code}{number} format")
	}

	eaterID, err := s.eaterSvc.SignupEater(ctx, req.GetPhoneNumber())
	if err != nil {
		return nil, err
	}

	return &pb.SignupEaterResponse{
		EaterId: eaterID,
	}, nil
}

func (s *eaterAppSvcImpl) ConfirmSMSCode(ctx context.Context, req *pb.ConfirmSmsCodeRequest) (*pb.ConfirmSmsCodeResponse, error) {
	if req.GetEaterId() == "" {
		return nil, errors.New("invalid or missing eaterId")
	}

	if req.GetSmsCode() == "" {
		return nil, errors.New("invalid or missing code")
	}

	profile, err := s.eaterSvc.ConfirmSMSCode(ctx, req.GetEaterId(), req.GetSmsCode())
	if err != nil {
		return nil, err
	}

	token, err := s.tokenSvc.CreateToken(ctx, profile.EaterID)
	if err != nil {
		return nil, err
	}

	response := pb.ConfirmSmsCodeResponse{
		Profile: dtos.ToEaterProfilePB(profile),
		Token:   token,
	}
	return &response, nil
}

func (s *eaterAppSvcImpl) GetEaterProfile(ctx context.Context, req *pb.GetEaterProfileRequest) (*pb.GetEaterProfileResponse, error) {
	if req.GetEaterId() == "" {
		return nil, errors.New("invalid or missing eaterId")
	}

	profile, err := s.eaterSvc.GetEaterProfile(ctx, req.GetEaterId())
	if err != nil {
		return nil, err
	}

	return &pb.GetEaterProfileResponse{
		Profile: dtos.ToEaterProfilePB(profile),
	}, nil
}

func (s *eaterAppSvcImpl) UpdateEaterProfile(ctx context.Context, req *pb.UpdateEaterProfileRequest) (*pb.UpdateEaterProfileResponse, error) {
	if req.GetName() == "" {
		return nil, fmt.Errorf("Invalid or mising profile name: %s", req.GetName())
	}

	profile, err := s.eaterSvc.UpdateEaterProfile(ctx, req.GetEaterId(), req.GetName(), req.GetImageUrl())
	if err != nil {
		return nil, err
	}

	return &pb.UpdateEaterProfileResponse{
		Profile: dtos.ToEaterProfilePB(profile),
	}, nil

}
