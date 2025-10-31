package usecase

import (
	"context"
	"test-interview-kc/internal/dto/request"
	"test-interview-kc/internal/dto/response"
)

type WalletWithdrawalUseCase interface {
	Withdraw(ctx context.Context, req request.WalletWithdrawalRequest) error
}

type WalletAccountUseCase interface {
	GetAccountDetails(ctx context.Context, accountID string) (response.WalletAccountDetailResponse, error)
}
