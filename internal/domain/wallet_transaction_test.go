package domain

import (
	"testing"
)

func TestCreateWalletTransaction(t *testing.T) {
	tests := []struct {
		name            string
		walletAccountID string
		requestID       string
		amount          float64
		txType          string
		entryType       string
		description     string
		author          string
		wantErr         bool
	}{
		{
			name:            "valid transaction",
			walletAccountID: "wallet1",
			requestID:       "req1",
			amount:          100.5,
			txType:          "DEBIT",
			entryType:       "WITHDRAWAL",
			description:     "test withdrawal",
			author:          "system",
			wantErr:         false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			trx, err := CreateWalletTransaction(tt.walletAccountID, tt.requestID, tt.amount, tt.txType, tt.entryType, tt.description, tt.author)
			if (err != nil) != tt.wantErr {
				t.Fatalf("expected error: %v, got: %v", tt.wantErr, err)
			}
			if err == nil {
				if trx.WalletAccountID != tt.walletAccountID {
					t.Errorf("expected WalletAccountID %v, got %v", tt.walletAccountID, trx.WalletAccountID)
				}
				if trx.RequestID != tt.requestID {
					t.Errorf("expected RequestID %v, got %v", tt.requestID, trx.RequestID)
				}
				if trx.Amount != tt.amount {
					t.Errorf("expected Amount %v, got %v", tt.amount, trx.Amount)
				}
				if trx.Type != tt.txType {
					t.Errorf("expected Type %v, got %v", tt.txType, trx.Type)
				}
				if trx.EntryType != tt.entryType {
					t.Errorf("expected EntryType %v, got %v", tt.entryType, trx.EntryType)
				}
				if trx.Description != tt.description {
					t.Errorf("expected Description %v, got %v", tt.description, trx.Description)
				}
				if trx.CreatedBy != tt.author {
					t.Errorf("expected CreatedBy %v, got %v", tt.author, trx.CreatedBy)
				}
				if trx.UpdatedBy != tt.author {
					t.Errorf("expected UpdatedBy %v, got %v", tt.author, trx.UpdatedBy)
				}
				if trx.TransactionDate == nil || trx.CreatedAt == nil || trx.UpdatedAt == nil {
					t.Errorf("expected non-nil timestamps")
				}
				if trx.ID == "" {
					t.Errorf("expected non-empty ID")
				}
			}
		})
	}
}
