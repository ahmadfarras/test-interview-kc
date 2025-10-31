package repository

import (
	"context"
	"test-interview-kc/internal/domain"
	"test-interview-kc/pkg/logger"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type walletTransactionRepositoryImpl struct {
	DB  *gorm.DB
	log *zap.Logger
}

// Create implements WalletTransactionRepository.
func (w *walletTransactionRepositoryImpl) Create(ctx context.Context, tx *domain.WalletTransaction) error {
	log := logger.FromContext(ctx, w.log)
	result := w.DB.WithContext(ctx).Create(tx)
	if err := result.Error; err != nil {
		log.Error("Failed to create wallet transaction", zap.Error(err))
		return err
	}
	return nil
}

// IsAlreadyProcessed implements WalletTransactionRepository.
func (w *walletTransactionRepositoryImpl) IsAlreadyProcessed(ctx context.Context, requestID string) (bool, error) {
	log := logger.FromContext(ctx, w.log)

	var count int64
	err := w.DB.WithContext(ctx).Model(&domain.WalletTransaction{}).Where("request_id = ?", requestID).Count(&count).Error
	if err != nil {
		log.Error("Failed to check if transaction is already processed", zap.Error(err))
		return false, err
	}
	return count > 0, nil
}

func NewWalletTransactionRepository(db *gorm.DB, log *zap.Logger) WalletTransactionRepository {
	return &walletTransactionRepositoryImpl{DB: db, log: log}
}
