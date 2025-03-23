package model

// User is a struct that represent a user
type User struct {
	ID      string  `json:"id"`
	Name    string  `json:"name"`
	Balance float64 `json:"balance"`
}

// Transaction is a struct that represent a transaction
type Transaction struct {
	ID     string  `json:"id"`
	UserID string  `json:"user_id"`
	Type   string  `json:"type"`
	Amount float64 `json:"amount"`
}

// TransactionTypeDeposit is a constant that represent a deposit transaction type
const TransactionTypeDeposit = "deposit"

// TransactionTypeWithdraw is a constant that represent a withdraw transaction type
const TransactionTypeWithdraw = "withdraw"

// TransactionTypeInitialBalance is a constant that represent a initial balance transaction type
const TransactionTypeInitialBalance = "initial_balance"
