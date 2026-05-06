package services

import (
	"context"
	"errors"

	dtos "github.com/javohir01/eater-service/src/application/dtos/eater"
	eatersvc "github.com/javohir01/eater-service/src/domain/eater/services"
	"github.com/javohir01/eater-service/src/infrastructure/jwt"
	"github.com/javohir01/eater-service/src/infrastructure/validator"
)

type EaterApplicationService interface {
	SignupEater(ctx context.Context, phoneNumber string) (*dtos.EaterSignupResponse, error)
	ConfirmSMSCode(ctx context.Context, eaterID, smsCode string) (*dtos.ConfirmSmsCodeResponse, error)
	UpdateEaterProfile(ctx context.Context, eaterID, name, imageUrl string) (*dtos.UpdateEaterProfileResponse, error)
	GetEaterProfile(ctx context.Context, eaterID string) (*dtos.GetEaterProfileResponse, error)
}

type eaterAppSvcImpl struct {
	eaterSvc eatersvc.EaterService
	tokenSvc jwt.Service
}

func NewEaterApplicationService(eaterSvc eatersvc.EaterService, tokenSvc jwt.Service) EaterApplicationService {
	return &eaterAppSvcImpl{
		eaterSvc: eaterSvc,
		tokenSvc: tokenSvc,
	}
}

func (s *eaterAppSvcImpl) SignupEater(ctx context.Context, phoneNumber string) (*dtos.EaterSignupResponse, error) {
	if !validator.ValidateUzPhoneNumber(phoneNumber) {
		return nil, errors.New("invalid phone number")
	}

	eaterID, err := s.eaterSvc.SignupEater(ctx, phoneNumber)

	if err != nil {
		return nil, err
	}

	return dtos.NewEaterSignupResponse(eaterID), nil
}

func (s *eaterAppSvcImpl) ConfirmSMSCode(ctx context.Context, eaterID, smsCode string) (*dtos.ConfirmSmsCodeResponse, error) {
	if eaterID == "" || smsCode == "" {
		return nil, errors.New("eater ID and SMS code cannot be empty")
	}

	profile, err := s.eaterSvc.ConfirmSMSCode(ctx, eaterID, smsCode)
	if err != nil {
		return nil, err
	}

	token, err := s.tokenSvc.CreateToken(ctx, profile.EaterID)
	if err != nil {
		return nil, err
	}

	response := dtos.ConfirmSmsCodeResponse{
		Profile: profile,
		Token:   token,
	}
	return &response, nil
}

func (s *eaterAppSvcImpl) UpdateEaterProfile(ctx context.Context, eaterID, name, imageUrl string) (*dtos.UpdateEaterProfileResponse, error) {
	if name == "" {
		return nil, errors.New("name cannot be empty")
	}

	profile, err := s.eaterSvc.UpdateEaterProfile(ctx, eaterID, name, imageUrl)
	if err != nil {
		return nil, err
	}
	return dtos.NewUpdateEaterProfileResponse(profile), nil
}

func (s *eaterAppSvcImpl) GetEaterProfile(ctx context.Context, eaterID string) (*dtos.GetEaterProfileResponse, error) {
	if eaterID == "" {
		return nil, errors.New("invalid eater ID")
	}

	profile, err := s.eaterSvc.GetEaterProfile(ctx, eaterID)
	if err != nil {
		return nil, err
	}
	return dtos.NewGetEaterProfileResponse(profile), nil
}
