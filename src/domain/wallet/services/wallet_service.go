package services

import (
	"context"
	"time"

	"github.com/javohir01/eater-service/src/domain/wallet/models"
	"github.com/javohir01/eater-service/src/domain/wallet/repositories"
)

type WalletService interface {
	CreateWallet(ctx context.Context, ID, EaterID, Number string, IsVerifed bool) *models.PaymentCard
	GetWalletByCardId(ctx context.Context, id string) (*models.PaymentCard, error)
	GetWalletsListByEaterId(ctx context.Context, eater_id string) ([]*models.PaymentCard, error)
	DeleteWallet(ctx context.Context, id string) error
}

type walletSrvImpl struct {
	walletRepo repositories.WalletRepository
}

func NewWalletService(
	walletRepo repositories.WalletRepository,
) WalletService {
	return &walletSrvImpl{
		walletRepo: walletRepo,
	}
}

func (s *walletSrvImpl) CreateWallet(ctx context.Context, ID, EaterID, Number string, IsVerifed bool) *models.PaymentCard {
	wallet := &models.PaymentCard{
		ID:        ID,
		EaterID:   EaterID,
		Number:    Number,
		IsVerifed: IsVerifed,
		CreatedAt: time.Now().UTC(),
	}

	err := s.walletRepo.CreateWallet(ctx, wallet)
	if err != nil {
		return nil
	}

	return wallet
}

func (s *walletSrvImpl) GetWalletByCardId(ctx context.Context, id string) (*models.PaymentCard, error) {
	wallet, err := s.walletRepo.GetWalletByCardId(ctx, id)

	if err != nil {
		return nil, err
	}
	return wallet, nil
}

func (s *walletSrvImpl) GetWalletsListByEaterId(ctx context.Context, eater_id string) ([]*models.PaymentCard, error) {
	wallets, err := s.walletRepo.GetWalletsListByEaterId(ctx, eater_id)

	if err != nil {
		return nil, err
	}
	return wallets, nil
}

func (s *walletSrvImpl) DeleteWallet(ctx context.Context, id string) error {
	err := s.walletRepo.DeleteWallet(ctx, id)

	if err != nil {
		return err
	}

	return nil
}
