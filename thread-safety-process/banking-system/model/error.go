package model

import "fmt"

// ErrInsufficientBalance will indicates that the balance is not enough to withdraw
var ErrInsufficientBalance = fmt.Errorf("insufficient balance")

// ErrAccountLocked will indicates that the account is locked by other transaction
var ErrAccountLocked = fmt.Errorf("account is locked by other transaction")

// ErrInvalidTransactionType will indicates that the transaction type is invalid
var ErrInvalidTransactionType = fmt.Errorf("invalid transaction type")

// ErrInvalidTransactionAmount will indicates that the transaction amount is invalid
var ErrInvalidTransactionAmount = fmt.Errorf("invalid transaction amount")
