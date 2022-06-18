package repository

import (
	"fmt"

	test "github.com/Olexander753/constanta_test"
	"github.com/jmoiron/sqlx"
)

type TransactionsListPostgres struct {
	db *sqlx.DB
}

func NewTransactionsListPostgres(db *sqlx.DB) *TransactionsListPostgres {
	return &TransactionsListPostgres{db: db}
}

func (r *TransactionsListPostgres) GetListByID(user_id string) ([]test.Transactions, error) {
	var list []test.Transactions

	query := fmt.Sprintf("SELECT *FROM %s WHERE user_id = $1", transactionsTable)
	err := r.db.Select(&list, query, user_id)

	return list, err
}

func (r *TransactionsListPostgres) GetListByEmail(email string) ([]test.Transactions, error) {
	var list []test.Transactions

	query := fmt.Sprintf("SELECT transactions.id, transactions.user_id, transactions.sum, transactions.currensy, transactions.date_create, transactions.date_last_update, transactions.status FROM %s inner join %s on users.id = transactions.user_id WHERE email = $1", transactionsTable, usersTable)
	err := r.db.Select(&list, query, email)

	return list, err
}
