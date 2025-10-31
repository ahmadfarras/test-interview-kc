package error

import (
	"errors"
)

// Wallet Withdrawal error codes
var (
	// ErrInsufficientFunds indicates that the wallet has insufficient funds for the requested operation.
	ErrInsufficientFunds  = errors.New("insufficient funds")
	ErrIsAlreadyProcessed = errors.New("transaction already processed")
)
