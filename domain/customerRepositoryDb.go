package domain

import (
	"database/sql"

	"github.com/ShaizanKhan/go-banking-lib/errs"
	"github.com/ShaizanKhan/go-banking-lib/logger"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type CustomRepositoryDb struct {
	client *sqlx.DB
}

func (d CustomRepositoryDb) FindAll(status string) ([]Customer, *errs.AppError) {
	var err error
	customers := make([]Customer, 0)
	if status == "" {
		findAllSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers"
		err = d.client.Select(&customers, findAllSql)
	} else {
		findAllSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers where status = ?"
		err = d.client.Select(&customers, findAllSql, status)
	}

	if err != nil {
		logger.Error("Error while scanning customers table " + err.Error())
		return nil, errs.NewUnExpectedError("unexpected database error")
	}

	return customers, nil
}

func (d CustomRepositoryDb) ById(id string) (*Customer, *errs.AppError) {
	customerSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers where customer_id = ?"

	var c Customer
	err := d.client.Get(&c, customerSql, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("Customer not found")
		} else {
			logger.Error("Error while scanning customers table " + err.Error())
			return nil, errs.NewUnExpectedError("unexpected database error")
		}
	}
	return &c, nil
}

func NewCustomerRepositoryDb(dbClient *sqlx.DB) CustomRepositoryDb {
	return CustomRepositoryDb{client: dbClient}
}
