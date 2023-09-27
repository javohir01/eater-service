package wallet

import (
	"context"

	"github.com/javohir01/eater-service/src/domain/wallet/models"
	"github.com/javohir01/eater-service/src/domain/wallet/repositories"
	"gorm.io/gorm"
)

const (
	tableWallet = "eater.wallet"
)

type walletRepoImpl struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) repositories.WalletRepository {
	return &walletRepoImpl{
		db: db,
	}
}

func (r *walletRepoImpl) WithContext(ctx context.Context, f func(r repositories.WalletRepository) error) error {
	if err := r.db.Transaction(func(tx *gorm.DB) error {
		r := walletRepoImpl{
			db: tx,
		}
		if err := f(&r); err != nil {
			return err
		}
		return nil
	}); err != nil {
		return err
	}
	return nil
}

func (r *walletRepoImpl) CreateWallet(ctx context.Context, wallet *models.PaymentCard) error {
	result := r.db.WithContext(ctx).Table(tableWallet).Create(wallet)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *walletRepoImpl) DeleteWallet(ctx context.Context, id string) error {
	var wallet models.PaymentCard
	result := r.db.WithContext(ctx).Table(tableWallet).Delete(&wallet, "id = ?", id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *walletRepoImpl) GetWalletByCardId(ctx context.Context, card_id string) (*models.PaymentCard, error) {
	var wallet models.PaymentCard
	result := r.db.WithContext(ctx).Table(tableWallet).First(&wallet, "id = ?", card_id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &wallet, nil
}

func (r *walletRepoImpl) GetWalletsListByEaterId(ctx context.Context, eater_id string) ([]*models.PaymentCard, error) {
	var wallet []*models.PaymentCard
	result := r.db.WithContext(ctx).Table(tableWallet).Where(&wallet, "id = ?", eater_id)
	if result.Error != nil {
		return nil, result.Error
	}
	return wallet, nil
}
