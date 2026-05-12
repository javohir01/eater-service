package services

import (
	"context"
	"errors"

	"github.com/javohir01/eater-service/src/application/dtos/wallet"
	"github.com/javohir01/eater-service/src/domain/wallet/models"
	walletsvc "github.com/javohir01/eater-service/src/domain/wallet/services"
)

type WalletApplicationService interface {
	AddCard(ctx context.Context, eaterID string, req *wallet.AddCardRequest) (*wallet.CardResponse, error)
	GetCard(ctx context.Context, cardID string) (*wallet.CardResponse, error)
	DeleteCard(ctx context.Context, cardID string) error
	GetCardsByEaterID(ctx context.Context, eaterID string) ([]*wallet.CardResponse, error)
}

type walletAppSvcImpl struct {
	walletSvc walletsvc.WalletService
}

func NewWalletApplicationService(walletSvc walletsvc.WalletService) WalletApplicationService {
	return &walletAppSvcImpl{
		walletSvc: walletSvc,
	}
}

func (s *walletAppSvcImpl) AddCard(ctx context.Context, eaterID string, req *wallet.AddCardRequest) (*wallet.CardResponse, error) {
	if req == nil {
		return nil, errors.New("add card request cannot be nil")
	}

	if eaterID == "" {
		return nil, errors.New("eater ID cannot be empty")
	}

	if req.Number == "" {
		return nil, errors.New("card number cannot be empty")
	}

	card, err := s.walletSvc.AddCard(ctx, eaterID, req.Number, req.CardToken)
	if err != nil {
		return nil, err
	}

	return mapToCardResponse(card), nil
}

func (s *walletAppSvcImpl) GetCard(ctx context.Context, cardID string) (*wallet.CardResponse, error) {
	if cardID == "" {
		return nil, errors.New("card ID cannot be empty")
	}

	card, err := s.walletSvc.GetCard(ctx, cardID)
	if err != nil {
		return nil, err
	}

	if card == nil {
		return nil, errors.New("card not found")
	}

	return mapToCardResponse(card), nil
}

func (s *walletAppSvcImpl) DeleteCard(ctx context.Context, cardID string) error {
	if cardID == "" {
		return errors.New("card ID cannot be empty")
	}

	return s.walletSvc.DeleteCard(ctx, cardID)
}

func (s *walletAppSvcImpl) GetCardsByEaterID(ctx context.Context, eaterID string) ([]*wallet.CardResponse, error) {
	if eaterID == "" {
		return nil, errors.New("eater ID cannot be empty")
	}

	cards, err := s.walletSvc.GetCardsByEaterID(ctx, eaterID)
	if err != nil {
		return nil, err
	}

	var responses []*wallet.CardResponse
	for _, card := range cards {
		responses = append(responses, mapToCardResponse(card))
	}

	return responses, nil
}

func mapToCardResponse(card *models.PaymentCard) *wallet.CardResponse {
	return &wallet.CardResponse{
		ID:         card.ID,
		EaterID:    card.EaterID,
		Number:     card.Number,
		IsVerified: card.IsVerified,
		CreatedAt:  card.CreatedAt,
		UpdatedAt:  card.UpdatedAt,
	}
}
