package response

import (
	"test-interview-kc/internal/domain"
	"time"
)

type WalletAccountDetailResponse struct {
	ID        string     `json:"id"`
	Name      string     `json:"name"`
	Balance   float64    `json:"balance"`
	CreatedAt *time.Time `json:"created_at"`
	CreatedBy string     `json:"created_by"`
	UpdatedAt *time.Time `json:"updated_at"`
	UpdatedBy string     `json:"updated_by"`
}

func ToWalletAccountDetailResponse(
	account domain.WalletAccount,
) WalletAccountDetailResponse {
	return WalletAccountDetailResponse{
		ID:        account.ID,
		Name:      account.Name,
		Balance:   account.Balance,
		CreatedAt: account.CreatedAt,
		CreatedBy: account.CreatedBy,
		UpdatedAt: account.UpdatedAt,
		UpdatedBy: account.UpdatedBy,
	}
}
