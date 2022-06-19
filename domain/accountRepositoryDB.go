package domain

import (
	"ashishi-banking/errs"
	"ashishi-banking/logger"
	"github.com/jmoiron/sqlx"
	"strconv"
)

type AccountRepositoryDb struct {
	client *sqlx.DB
}

func NewAccountRepositoryDb(dbClinet *sqlx.DB) AccountRepositoryDb {
	return AccountRepositoryDb{dbClinet}
}

func (d AccountRepositoryDb) Save(a Account) (*Account, *errs.AppError) {
	sqlInsert := "INSERT INTO accounts (customer_id, opening_id, account_type, amount, status) value (?, ?, ?, ?, ?)"
	result, err := d.client.Exec(sqlInsert, a.CustomerId, a.OpeningDate, a.AccountType, a.Amount, a.Status)
	if err != nil {
		logger.Error("Error where creating new account: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error form database")
	}

	id, err := result.LastInsertId()
	if err != nil {
		logger.Error("Error while getting last insert id: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error form database")
	}

	a.AccountId = strconv.FormatInt(id, 10)

	return &a, nil
}
