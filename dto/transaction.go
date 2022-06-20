package dto

import (
	"ashishi-banking/domain"
	"ashishi-banking/errs"
)

const WITHDRAWAL = "withdrawal"
const DEPOSIT = "deposit"

type TransactionRequest struct {
	AccountId       string  `json:"account_id"`
	Amount          float64 `json:"amount"`
	TransactionType string  `json:"transaction_type"`
	TransactionDate string  `json:"transaction_date"`
	CustomerId      string  `json:"-"`
}

func (t TransactionRequest) IsTransactionTypeWithdrawal() bool {
	if t.TransactionType == domain.WITHDRAWAL {
		return true
	}
	return false
}

func (t TransactionRequest) Validate() *errs.AppError {
	if t.TransactionType != WITHDRAWAL && t.TransactionType != DEPOSIT {
		return errs.NewValidationError("Transactoin type can only be deposit or withdrawal")
	}

	if t.Amount < 0 {
		return errs.NewValidationError("Amount cannot be less than zero")
	}

	return nil
}

type TransactionResponse struct {
	TransactionId   string  `json:"transaction_id"`
	AccountId       string  `json:"account_id"`
	Amount          float64 `json:"new_balance"`
	TransactionType string  `json:"transaction_type"`
	TransactionDate string  `json:"transaction_date"`
}
