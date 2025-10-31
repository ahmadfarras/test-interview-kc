package repository

import (
	"context"
	"test-interview-kc/internal/domain"
	"test-interview-kc/pkg/logger"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type walletAccountRepositoryImpl struct {
	DB  *gorm.DB
	log *zap.Logger
}

// GetByID implements WalletAccountRepository.
func (w *walletAccountRepositoryImpl) GetByID(ctx context.Context, accountID string) (*domain.WalletAccount, error) {
	log := logger.FromContext(ctx, w.log)

	var account domain.WalletAccount
	err := w.DB.WithContext(ctx).First(&account, "id = ?", accountID).Error
	if err != nil {
		log.Error("Failed to get wallet account", zap.Error(err))
		return nil, err
	}
	return &account, nil
}

// Update implements WalletAccountRepository.
func (w *walletAccountRepositoryImpl) Update(ctx context.Context, account *domain.WalletAccount) error {
	log := logger.FromContext(ctx, w.log)

	if err := w.DB.WithContext(ctx).Save(account).Error; err != nil {
		log.Error("Failed to update wallet account", zap.Error(err))
		return err
	}

	return nil
}

func NewWalletAccountRepository(db *gorm.DB, log *zap.Logger) WalletAccountRepository {
	return &walletAccountRepositoryImpl{DB: db, log: log}
}
