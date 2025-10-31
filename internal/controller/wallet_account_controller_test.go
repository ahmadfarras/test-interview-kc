package controller

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"test-interview-kc/internal/dto/response"
	"test-interview-kc/internal/usecase/mocks"

	"github.com/gofiber/fiber/v2"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func TestWalletAccountController_GetAccountDetails(t *testing.T) {
	app := fiber.New()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUseCase := mocks.NewMockWalletAccountUseCase(ctrl)
	controller := NewWalletAccountController(mockUseCase, zap.NewNop())
	app.Get("/wallets/:id", controller.GetAccountDetails)

	accountResp := response.WalletAccountDetailResponse{ID: "1", Name: "Test", Balance: 1000}

	tests := []struct {
		name       string
		url        string
		param      string
		mockSetup  func()
		wantStatus int
	}{
		{
			name:  "success",
			url:   "/wallets/1",
			param: "1",
			mockSetup: func() {
				mockUseCase.EXPECT().GetAccountDetails(gomock.Any(), "1").Return(accountResp, nil)
			},
			wantStatus: http.StatusOK,
		},
		{
			name:       "missing id",
			url:        "/wallets/",
			param:      "",
			mockSetup:  func() {},
			wantStatus: http.StatusNotFound, // Fiber returns 404 for missing param
		},
		{
			name:  "not found",
			url:   "/wallets/2",
			param: "2",
			mockSetup: func() {
				mockUseCase.EXPECT().GetAccountDetails(gomock.Any(), "2").Return(response.WalletAccountDetailResponse{}, gorm.ErrRecordNotFound)
			},
			wantStatus: http.StatusNotFound,
		},
		{
			name:  "internal error",
			url:   "/wallets/3",
			param: "3",
			mockSetup: func() {
				mockUseCase.EXPECT().GetAccountDetails(gomock.Any(), "3").Return(response.WalletAccountDetailResponse{}, errors.New("db error"))
			},
			wantStatus: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockSetup()
			req := httptest.NewRequest(http.MethodGet, tt.url, nil)
			resp, _ := app.Test(req)
			assert.Equal(t, tt.wantStatus, resp.StatusCode)
		})
	}
}
