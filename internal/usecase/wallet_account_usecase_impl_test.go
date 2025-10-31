package usecase

import (
	"context"
	"errors"
	"testing"

	"test-interview-kc/internal/domain"
	"test-interview-kc/internal/dto/response"
	"test-interview-kc/internal/repository/mocks"

	"go.uber.org/zap"
	"github.com/golang/mock/gomock"
)

func TestWalletAccountUseCaseImpl_GetAccountDetails(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockWalletAccountRepository(ctrl)
	uc := &walletAccountUseCaseImpl{
		walletAccountRepo: mockRepo,
		log:               zap.NewNop(),
	}

	account := &domain.WalletAccount{ID: "1", Name: "Test", Balance: 1000}
	resp := response.ToWalletAccountDetailResponse(*account)

	testCases := []struct {
		name      string
		setupMock func()
		want      response.WalletAccountDetailResponse
		wantErr   bool
	}{
		{
			name: "success",
			setupMock: func() {
				mockRepo.EXPECT().GetByID(gomock.Any(), "1").Return(account, nil)
			},
			want:    resp,
			wantErr: false,
		},
		{
			name: "repo error",
			setupMock: func() {
				mockRepo.EXPECT().GetByID(gomock.Any(), "2").Return(nil, errors.New("not found"))
			},
			want:    response.WalletAccountDetailResponse{},
			wantErr: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.setupMock()
			got, err := uc.GetAccountDetails(context.Background(), func() string {
				if tc.name == "success" {
					return "1"
				}
				return "2"
			}())
			if (err != nil) != tc.wantErr {
				t.Fatalf("expected error: %v, got: %v", tc.wantErr, err)
			}
			if got != tc.want {
				t.Errorf("expected: %+v, got: %+v", tc.want, got)
			}
		})
	}
}
