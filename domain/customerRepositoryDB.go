package domain

import (
	"ashishi-banking/errs"
	"ashishi-banking/logger"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"time"
)

type CustomerRepositoryDB struct {
	client *sqlx.DB
}

func (d CustomerRepositoryDB) FindAll(status string) ([]Customer, *errs.AppError) {
	var err error
	customers := make([]Customer, 0)

	if status == "" {
		findAllSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers"
		err = d.client.Select(&customers, findAllSql)
		//rows, err := d.client.Query(findAllSql)
	} else {
		findAllSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers where status = ?"
		err = d.client.Select(&customers, findAllSql, status)
	}

	if err != nil {
		logger.Error("Error while querying customer table" + err.Error())
		return nil, errs.NewUnexpectedError("Customer not found")
	}

	//err = sqlx.StructScan(rows, &customers)
	//if err != nil {
	//	logger.Error("Error while scanning customer " + err.Error())
	//	return nil, errs.NewNotFoundError("Customer not found")
	//}

	//for rows.Next() {
	//	var c Customer
	//	err := rows.Scan(&c.Id, &c.Name, &c.City, &c.Zipcode, &c.Status, &c.DateofBirth)
	//	if err != nil {
	//		logger.Error("Error while scanning customer " + err.Error())
	//		return nil, err
	//	}
	//
	//	customers = append(customers, c)
	//}

	return customers, nil
}

func (d CustomerRepositoryDB) ById(id string) (*Customer, *errs.AppError) {
	customerSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers where customer_id = ?"

	//row := d.client.QueryRow(customerSql, id)
	var c Customer
	err := d.client.Get(&c, customerSql, id)
	//err := row.Scan(&c.Id, &c.Name, &c.City, &c.Zipcode, &c.Status, &c.DateofBirth)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("Customer not found")
		} else {
			logger.Error("Error while scanning customer " + err.Error())
			return nil, errs.NewUnexpectedError("Unexpected database error")
		}
	}
	return &c, nil
}

func NewCustomerRepositoryDb() CustomerRepositoryDB {
	client, err := sqlx.Open("mysql", "root:rootadmin@tcp(localhost:3306)/banking")
	if err != nil {
		panic(err)
	}

	client.Ping()
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)
	return CustomerRepositoryDB{client}
}
