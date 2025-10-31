package usecase

import (
	"context"
	"errors"
	"testing"

	"test-interview-kc/internal/domain"
	"test-interview-kc/internal/dto/request"
	"test-interview-kc/internal/repository/mocks"

	"github.com/golang/mock/gomock"
	"go.uber.org/zap"
)

func TestWalletWithdrawalUseCase_Withdraw(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	cases := []struct {
		name                string
		setupAccount        func() *domain.WalletAccount
		getByIDErr          error
		alreadyProcessed    bool
		alreadyProcessedErr error
		canWithdraw         bool
		createTxErr         error
		updateErr           error
		expectErr           error
	}{
		{
			name: "success",
			setupAccount: func() *domain.WalletAccount {
				return &domain.WalletAccount{ID: "1", Balance: 1000}
			},
			canWithdraw: true,
		},
		{
			name:       "account not found",
			getByIDErr: errors.New("not found"),
			expectErr:  errors.New("not found"),
		},
		{
			name: "already processed",
			setupAccount: func() *domain.WalletAccount {
				return &domain.WalletAccount{ID: "1", Balance: 1000}
			},
			alreadyProcessed: true,
			expectErr:        errors.New("already processed"),
		},
		{
			name: "insufficient funds",
			setupAccount: func() *domain.WalletAccount {
				return &domain.WalletAccount{ID: "1", Balance: 1000}
			},
			canWithdraw: false,
			expectErr:   errors.New("insufficient funds"),
		},
		{
			name: "create tx error",
			setupAccount: func() *domain.WalletAccount {
				return &domain.WalletAccount{ID: "1", Balance: 1000}
			},
			canWithdraw: true,
			createTxErr: errors.New("create tx error"),
			expectErr:   errors.New("create tx error"),
		},
		{
			name: "update error",
			setupAccount: func() *domain.WalletAccount {
				return &domain.WalletAccount{ID: "1", Balance: 1000}
			},
			canWithdraw: true,
			updateErr:   errors.New("update error"),
			expectErr:   errors.New("update error"),
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			mockAccRepo := mocks.NewMockWalletAccountRepository(ctrl)
			mockTxRepo := mocks.NewMockWalletTransactionRepository(ctrl)

			if tc.getByIDErr != nil {
				mockAccRepo.EXPECT().GetByID(gomock.Any(), gomock.Any()).Return(nil, tc.getByIDErr)
			} else {
				acc := tc.setupAccount()
				if !tc.canWithdraw {
					acc.Balance = 0
				}
				mockAccRepo.EXPECT().GetByID(gomock.Any(), gomock.Any()).Return(acc, nil)
			}

			mockAccRepo.EXPECT().Update(gomock.Any(), gomock.Any()).DoAndReturn(
				func(ctx context.Context, acc *domain.WalletAccount) error {
					return tc.updateErr
				},
			).AnyTimes()

			mockTxRepo.EXPECT().IsAlreadyProcessed(gomock.Any(), gomock.Any()).Return(tc.alreadyProcessed, tc.alreadyProcessedErr).AnyTimes()
			mockTxRepo.EXPECT().Create(gomock.Any(), gomock.Any()).DoAndReturn(
				func(ctx context.Context, trx *domain.WalletTransaction) error {
					return tc.createTxErr
				},
			).AnyTimes()

			uc := &walletWithdrawalUseCaseImpl{
				walletAccountRepo: mockAccRepo,
				walletTxRepo:      mockTxRepo,
				log:               zap.NewNop(),
			}
			err := uc.Withdraw(context.Background(), request.WalletWithdrawalRequest{WalletID: "1", Amount: 100, RequestID: "req-1"})
			if (err != nil && tc.expectErr == nil) || (err == nil && tc.expectErr != nil) {
				t.Fatalf("expected error: %v, got: %v", tc.expectErr, err)
			}
		})
	}
}
