package domain

import (
	"testing"
	"time"
)

func TestWalletAccount_CanWithdraw(t *testing.T) {
	tests := []struct {
		name    string
		balance float64
		amount  float64
		want    bool
	}{
		{"sufficient balance", 100, 50, true},
		{"exact balance", 100, 100, true},
		{"insufficient balance", 100, 150, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			wa := &WalletAccount{Balance: tt.balance}
			got := wa.CanWithdraw(tt.amount)
			if got != tt.want {
				t.Errorf("CanWithdraw(%v) = %v, want %v", tt.amount, got, tt.want)
			}
		})
	}
}

func TestWalletAccount_Withdraw(t *testing.T) {
	tests := []struct {
		name     string
		startBal float64
		withdraw float64
		author   string
		wantBal  float64
		wantBy   string
	}{
		{"normal withdraw", 100, 30, "tester", 70, "tester"},
		{"zero withdraw", 100, 0, "nobody", 100, "nobody"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			wa := &WalletAccount{Balance: tt.startBal}
			wa.Withdraw(tt.withdraw, tt.author)
			if wa.Balance != tt.wantBal {
				t.Errorf("expected balance %v, got %v", tt.wantBal, wa.Balance)
			}
			if wa.UpdatedBy != tt.wantBy {
				t.Errorf("expected UpdatedBy '%v', got '%v'", tt.wantBy, wa.UpdatedBy)
			}
			if wa.UpdatedAt == nil || time.Since(*wa.UpdatedAt) > time.Second {
				t.Errorf("expected UpdatedAt to be set to now, got %v", wa.UpdatedAt)
			}
		})
	}
}

func TestWalletAccount_TableName(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{"table name", "wallet_account"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			wa := &WalletAccount{}
			if got := wa.TableName(); got != tt.want {
				t.Errorf("expected table name '%v', got '%v'", tt.want, got)
			}
		})
	}
}
