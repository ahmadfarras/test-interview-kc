package request

type WalletWithdrawalRequest struct {
	WalletID    string  `json:"wallet_id" validate:"required"`
	Amount      float64 `json:"amount" validate:"required,gt=0"`
	Description string  `json:"description" validate:"omitempty,max=255"`
	RequestID   string  `json:"request_id" validate:"required"`
}
