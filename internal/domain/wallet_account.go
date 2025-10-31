package domain

import "time"

type WalletAccount struct {
	ID        string     `gorm:"id"`
	Name      string     `gorm:"name"`
	Balance   float64    `gorm:"balance"`
	CreatedAt *time.Time `gorm:"created_at;autoCreateTime"`
	CreatedBy string     `gorm:"created_by"`
	UpdatedAt *time.Time `gorm:"updated_at;autoCreateTime;autoUpdateTime"`
	UpdatedBy string     `gorm:"updated_by"`
	DeletedAt *time.Time `gorm:"deleted_at"`
	DeletedBy string     `gorm:"deleted_by"`
}

func (w *WalletAccount) TableName() string {
	return "wallet_account"
}

func (w *WalletAccount) CanWithdraw(amount float64) bool {
	return w.Balance >= amount
}

func (w *WalletAccount) Withdraw(amount float64, author string) {
	w.Balance -= amount
	w.UpdatedAt = timePtr(time.Now())
	w.UpdatedBy = author
}
