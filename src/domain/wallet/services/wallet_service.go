package services

import (
	"context"

	"github.com/javohir01/eater-service/src/domain/wallet/models"
	"github.com/javohir01/eater-service/src/domain/wallet/repositories"
)

type WalletService interface {
	AddCard(ctx context.Context, eaterID, number, cardToken string) (*models.PaymentCard, error)
	GetCard(ctx context.Context, cardID string) (*models.PaymentCard, error)
	DeleteCard(ctx context.Context, cardID string) error
	GetCardsByEaterID(ctx context.Context, eaterID string) ([]*models.PaymentCard, error)
}

type walletSvcImpl struct {
	walletRepo repositories.WalletRepository
}

func NewWalletService(walletRepo repositories.WalletRepository) WalletService {
	return &walletSvcImpl{
		walletRepo: walletRepo,
	}
}

func (s *walletSvcImpl) AddCard(ctx context.Context, eaterID, number, cardToken string) (*models.PaymentCard, error) {
	return s.walletRepo.AddCard(ctx, eaterID, number, cardToken)
}

func (s *walletSvcImpl) GetCard(ctx context.Context, cardID string) (*models.PaymentCard, error) {
	return s.walletRepo.GetCard(ctx, cardID)
}

func (s *walletSvcImpl) DeleteCard(ctx context.Context, cardID string) error {
	return s.walletRepo.DeleteCard(ctx, cardID)
}

func (s *walletSvcImpl) GetCardsByEaterID(ctx context.Context, eaterID string) ([]*models.PaymentCard, error) {
	return s.walletRepo.GetCardsByEaterID(ctx, eaterID)
}
