package repositories

import (
	"context"

	"github.com/javohir01/eater-service/src/domain/wallet/models"
)

type WalletRepository interface {
	WithContext(ctx context.Context, f func(r WalletRepository) error) error
	CreateWallet(ctx context.Context, wallet *models.PaymentCard) error
	GetWalletByCardId(ctx context.Context, id string) (*models.PaymentCard, error)
	GetWalletsListByEaterId(ctx context.Context, eater_id string) ([]*models.PaymentCard, error)
	DeleteWallet(ctx context.Context, id string) error
}
