package domain

import (
	"ashishi-banking/dto"
	"ashishi-banking/errs"
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

type AccountRepository interface {
	Save(Account) (*Account, *errs.AppError)
	FindBy(accountId string) (*Account, *errs.AppError)
	SaveTransaction(transaction Transaction) (*Transaction, *errs.AppError)
}

func (a Account) CanWithdrawal(amount float64) bool {
	if a.Amount < amount {
		return false
	}
	return true
}
