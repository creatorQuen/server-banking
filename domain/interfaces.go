package domain

import "ashishi-banking/errs"

//go:generate mockgen -source=interfaces.go -destination=mocks.go -package=domain
type CustomerRepository interface {
	FindAll(string) ([]Customer, *errs.AppError)
	ById(string) (*Customer, *errs.AppError)
}

type AccountRepository interface {
	Save(Account) (*Account, *errs.AppError)
	FindBy(accountId string) (*Account, *errs.AppError)
	SaveTransaction(transaction Transaction) (*Transaction, *errs.AppError)
}
