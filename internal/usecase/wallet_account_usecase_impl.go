package usecase

import (
	"context"
	"test-interview-kc/internal/dto/response"
	"test-interview-kc/internal/repository"

	"go.uber.org/zap"
)

type walletAccountUseCaseImpl struct {
	walletAccountRepo repository.WalletAccountRepository
	log               *zap.Logger
}

// GetAccountDetails implements WalletAccountUseCase.
func (w *walletAccountUseCaseImpl) GetAccountDetails(
	ctx context.Context, accountID string,
) (response.WalletAccountDetailResponse, error) {
	account, err := w.walletAccountRepo.GetByID(ctx, accountID)
	if err != nil {
		return response.WalletAccountDetailResponse{}, err
	}

	return response.ToWalletAccountDetailResponse(*account), nil
}

func NewWalletAccountUseCase(
	walletAccountRepo repository.WalletAccountRepository,
	log *zap.Logger,
) WalletAccountUseCase {
	return &walletAccountUseCaseImpl{
		walletAccountRepo: walletAccountRepo,
		log:               log,
	}
}
