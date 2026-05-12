package wallet

import (
	"context"
	"time"

	"github.com/javohir01/eater-service/src/domain/wallet/models"
	"github.com/javohir01/eater-service/src/infrastructure/rand"
	"gorm.io/gorm"
)

const (
	TablePaymentCard = "eater.payment_card"
)

type WalletRepoImpl struct {
	db *gorm.DB
}

func NewWalletRepo(db *gorm.DB) *WalletRepoImpl {
	return &WalletRepoImpl{
		db: db,
	}
}

func (r *WalletRepoImpl) WithTx(ctx context.Context, f func(r *WalletRepoImpl) error) error {
	if err := r.db.Transaction(func(tx *gorm.DB) error {
		r := &WalletRepoImpl{
			db: tx,
		}
		if err := f(r); err != nil {
			return err
		}
		return nil
	}); err != nil {
		return err
	}
	return nil
}

func (r *WalletRepoImpl) AddCard(ctx context.Context, eaterID, number, cardToken string) (*models.PaymentCard, error) {
	card := &models.PaymentCard{
		ID:         rand.String(10),
		EaterID:    eaterID,
		Number:     number,
		CardToken:  cardToken,
		IsVerified: false,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	result := r.db.WithContext(ctx).Table(TablePaymentCard).Create(card)
	if result.Error != nil {
		return nil, result.Error
	}
	return card, nil
}

func (r *WalletRepoImpl) GetCard(ctx context.Context, cardID string) (*models.PaymentCard, error) {
	var card models.PaymentCard
	result := r.db.WithContext(ctx).Table(TablePaymentCard).First(&card, "id = ?", cardID)
	if result.Error != nil {
		return nil, result.Error
	}
	return &card, nil
}

func (r *WalletRepoImpl) DeleteCard(ctx context.Context, cardID string) error {
	result := r.db.WithContext(ctx).Table(TablePaymentCard).Where("id = ?", cardID).Delete(&models.PaymentCard{})
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *WalletRepoImpl) GetCardsByEaterID(ctx context.Context, eaterID string) ([]*models.PaymentCard, error) {
	var cards []*models.PaymentCard
	result := r.db.WithContext(ctx).Table(TablePaymentCard).Where("eater_id = ?", eaterID).Find(&cards)
	if result.Error != nil {
		return nil, result.Error
	}
	return cards, nil
}
