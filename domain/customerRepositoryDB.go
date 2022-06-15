package domain

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
)

type CustomerRepositoryDB struct {
	client *sql.DB
}

func (d CustomerRepositoryDB) FindAll() ([]Customer, error) {
	findAllSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers"

	rows, err := d.client.Query(findAllSql)
	if err != nil {
		log.Println("Error while querying customer table" + err.Error())
		return nil, err
	}

	customers := make([]Customer, 0)
	for rows.Next() {
		var c Customer
		err := rows.Scan(&c.Id, &c.Name, &c.City, &c.Zipcode, &c.Status, &c.DateofBirth)
		if err != nil {
			log.Println("Error while scanning customer " + err.Error())
			return nil, err
		}

		customers = append(customers, c)
	}

	return customers, nil
}

func NewCustomerRepositoryDb() CustomerRepositoryDB {
	client, err := sql.Open("mysql", "root:rootadmin@tcp(localhost:3306)/banking")
	if err != nil {
		panic(err)
	}

	client.Ping()
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)
	return CustomerRepositoryDB{client}
}
