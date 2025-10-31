package repository

import (
	"context"
	"test-interview-kc/internal/domain"
)

type WalletTransactionRepository interface {
	Create(ctx context.Context, tx *domain.WalletTransaction) error
	IsAlreadyProcessed(ctx context.Context, transactionID string) (bool, error)
}

type WalletAccountRepository interface {
	GetByID(ctx context.Context, accountID string) (*domain.WalletAccount, error)
	Update(ctx context.Context, account *domain.WalletAccount) error
}
