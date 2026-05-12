package repositories

import (
	"context"

	"github.com/javohir01/eater-service/src/domain/wallet/models"
)

type WalletRepository interface {
	AddCard(ctx context.Context, eaterID, number, cardToken string) (*models.PaymentCard, error)
	GetCard(ctx context.Context, cardID string) (*models.PaymentCard, error)
	DeleteCard(ctx context.Context, cardID string) error
	GetCardsByEaterID(ctx context.Context, eaterID string) ([]*models.PaymentCard, error)
}
