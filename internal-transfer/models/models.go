package models

type Account struct {
	AccountID int64   `json:"account_id"`
	Balance   float64 `json:"balance"`
}

type TransferRequest struct {
	FromAccountID int64   `json:"from_account_id"`
	ToAccountID   int64   `json:"to_account_id"`
	Amount        float64 `json:"amount"`
}
