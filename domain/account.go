package domain

import (
	"ashishi-banking/dto"
)

type Account struct {
	AccountId   string
	CustomerId  string
	OpeningDate string
	AccountType string
	Amount      float64
	Status      string
}

func (a Account) ToNewAccountResponseDto() dto.NewAccountResponse {
	return dto.NewAccountResponse{a.AccountId}
}

func (a Account) CanWithdrawal(amount float64) bool {
	if a.Amount < amount {
		return false
	}
	return true
}
