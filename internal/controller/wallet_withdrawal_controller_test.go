package controller

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"test-interview-kc/internal/error"
	"test-interview-kc/internal/usecase/mocks"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func TestWalletWithdrawalController_Withdraw(t *testing.T) {
	app := fiber.New()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUseCase := mocks.NewMockWalletWithdrawalUseCase(ctrl)
	val := validator.New()
	controller := NewWalletWithdrawalController(mockUseCase, val, zap.NewNop())
	app.Post("/wallets/:wallet_id/withdraw", controller.Withdraw)

	validBody := `{"amount":100}`

	tests := []struct {
		name         string
		walletID     string
		xRequestID   string
		body         string
		mockSetup    func()
		wantStatus   int
		wantContains string
	}{
		{
			name:       "success",
			walletID:   "1",
			xRequestID: "req-1",
			body:       validBody,
			mockSetup: func() {
				mockUseCase.EXPECT().Withdraw(gomock.Any(), gomock.Any()).Return(nil)
			},
			wantStatus:   http.StatusOK,
			wantContains: "withdrawal successful",
		},
		{
			name:         "missing X-Request-ID",
			walletID:     "1",
			xRequestID:   "",
			body:         validBody,
			mockSetup:    func() {},
			wantStatus:   http.StatusBadRequest,
			wantContains: "missing X-Request-ID",
		},
		{
			name:         "invalid body",
			walletID:     "1",
			xRequestID:   "req-1",
			body:         `invalid`,
			mockSetup:    func() {},
			wantStatus:   http.StatusBadRequest,
			wantContains: "invalid request body",
		},
		{
			name:         "validation failed",
			walletID:     "1",
			xRequestID:   "req-1",
			body:         `{}`,
			mockSetup:    func() {},
			wantStatus:   http.StatusBadRequest,
			wantContains: "VALIDATION_FAILED",
		},
		{
			name:       "wallet not found",
			walletID:   "1",
			xRequestID: "req-1",
			body:       validBody,
			mockSetup: func() {
				mockUseCase.EXPECT().Withdraw(gomock.Any(), gomock.Any()).Return(gorm.ErrRecordNotFound)
			},
			wantStatus:   http.StatusNotFound,
			wantContains: "WALLET_NOT_FOUND",
		},
		{
			name:       "insufficient funds",
			walletID:   "1",
			xRequestID: "req-1",
			body:       validBody,
			mockSetup: func() {
				mockUseCase.EXPECT().Withdraw(gomock.Any(), gomock.Any()).Return(error.ErrInsufficientFunds)
			},
			wantStatus:   http.StatusUnprocessableEntity,
			wantContains: "INSUFFICIENT_FUNDS",
		},
		{
			name:       "already processed",
			walletID:   "1",
			xRequestID: "req-1",
			body:       validBody,
			mockSetup: func() {
				mockUseCase.EXPECT().Withdraw(gomock.Any(), gomock.Any()).Return(error.ErrIsAlreadyProcessed)
			},
			wantStatus:   http.StatusUnprocessableEntity,
			wantContains: "ALREADY_PROCESSED",
		},
		{
			name:       "internal error",
			walletID:   "1",
			xRequestID: "req-1",
			body:       validBody,
			mockSetup: func() {
				mockUseCase.EXPECT().Withdraw(gomock.Any(), gomock.Any()).Return(errors.New("db error"))
			},
			wantStatus:   http.StatusInternalServerError,
			wantContains: "WITHDRAWAL_FAILED",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockSetup()
			req := httptest.NewRequest(http.MethodPost, "/wallets/"+tt.walletID+"/withdraw", strings.NewReader(tt.body))
			if tt.xRequestID != "" {
				req.Header.Set("X-Request-ID", tt.xRequestID)
			}
			resp, _ := app.Test(req)
			assert.Equal(t, tt.wantStatus, resp.StatusCode)
		})
	}
}
