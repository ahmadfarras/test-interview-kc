package domain

import (
	"time"

	"github.com/google/uuid"
)

type WalletTransaction struct {
	ID              string     `gorm:"id"`
	WalletAccountID string     `gorm:"wallet_account_id"`
	RequestID       string     `gorm:"request_id"`
	Amount          float64    `gorm:"amount"`
	Type            string     `gorm:"type"`       // e.g., "credit" or "debit"
	EntryType       string     `gorm:"entry_type"` // e.g., "withdrawal", "deposit", "transfer", "payment"
	Description     string     `gorm:"description"`
	TransactionDate *time.Time `gorm:"transaction_date"`
	CreatedAt       *time.Time `gorm:"created_at;autoCreateTime"`
	CreatedBy       string     `gorm:"created_by"`
	UpdatedAt       *time.Time `gorm:"updated_at;autoCreateTime;autoUpdateTime"`
	UpdatedBy       string     `gorm:"updated_by"`
	DeletedAt       *time.Time `gorm:"deleted_at"`
	DeletedBy       string     `gorm:"deleted_by"`
}

func (w *WalletTransaction) TableName() string {
	return "wallet_transaction"
}

func CreateWalletTransaction(
	walletAccountID, requestID string, amount float64,
	txType, entryType, description, author string,
) (*WalletTransaction, error) {
	id, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}

	return &WalletTransaction{
		ID:              id.String(),
		WalletAccountID: walletAccountID,
		RequestID:       requestID,
		Amount:          amount,
		Type:            txType,
		EntryType:       entryType,
		Description:     description,
		TransactionDate: timePtr(time.Now()),
		CreatedAt:       timePtr(time.Now()),
		CreatedBy:       author,
		UpdatedAt:       timePtr(time.Now()),
		UpdatedBy:       author,
	}, nil
}

func timePtr(t time.Time) *time.Time {
	return &t
}
