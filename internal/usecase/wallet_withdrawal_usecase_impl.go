package usecase

import (
	"context"
	"test-interview-kc/internal/domain"
	"test-interview-kc/internal/dto/request"
	"test-interview-kc/internal/repository"
	"test-interview-kc/pkg/logger"

	walletError "test-interview-kc/internal/error"

	"go.uber.org/zap"
)

type walletWithdrawalUseCaseImpl struct {
	walletAccountRepo repository.WalletAccountRepository
	walletTxRepo      repository.WalletTransactionRepository
	log               *zap.Logger
}

// Withdraw implements WalletWithdrawalUseCase.
func (w *walletWithdrawalUseCaseImpl) Withdraw(ctx context.Context, req request.WalletWithdrawalRequest) error {
	log := logger.FromContext(ctx, w.log)

	walletAccount, err := w.walletAccountRepo.GetByID(ctx, req.WalletID)
	if err != nil {
		log.Error("Failed to get wallet account", zap.Error(err))
		return err
	}

	isProcessed, err := w.walletTxRepo.IsAlreadyProcessed(ctx, req.RequestID)
	if err != nil {
		log.Error("Failed to check if transaction is already processed", zap.Error(err))
		return err
	}

	if isProcessed {
		log.Info("Withdrawal request already processed", zap.String("request_id", req.RequestID))
		return walletError.ErrIsAlreadyProcessed
	}

	if !walletAccount.CanWithdraw(req.Amount) {
		log.Info("Insufficient withdrawal amount", zap.Float64("requested_amount", req.Amount))
		return walletError.ErrInsufficientFunds
	}

	walletTrx, err := domain.CreateWalletTransaction(walletAccount.ID, req.RequestID, req.Amount, "DEBIT", "WITHDRAWAL", req.Description, "SYSTEM")
	if err != nil {
		log.Error("Failed to create wallet transaction", zap.Error(err))
		return err
	}

	err = w.walletTxRepo.Create(ctx, walletTrx)
	if err != nil {
		log.Error("Failed to save wallet transaction", zap.Error(err))
		return err
	}

	walletAccount.Withdraw(req.Amount, "SYSTEM")
	err = w.walletAccountRepo.Update(ctx, walletAccount)
	if err != nil {
		log.Error("Failed to update wallet account", zap.Error(err))
		return err
	}

	return nil
}

func NewWalletWithdrawalUseCase(
	walletAccountRepo repository.WalletAccountRepository,
	walletTxRepo repository.WalletTransactionRepository,
	log *zap.Logger,
) WalletWithdrawalUseCase {
	return &walletWithdrawalUseCaseImpl{
		walletAccountRepo: walletAccountRepo,
		walletTxRepo:      walletTxRepo,
		log:               log,
	}
}
