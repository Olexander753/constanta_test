package repository

import (
	"errors"
	"fmt"
	"math/rand"

	test "github.com/Olexander753/constanta_test"
	"github.com/jmoiron/sqlx"
)

type transactionItemPostgres struct {
	db *sqlx.DB
}

func NewTransactionItemPostgres(db *sqlx.DB) *transactionItemPostgres {
	return &transactionItemPostgres{db: db}
}

func (tran *transactionItemPostgres) CreateTransaction(transaction test.Transaction) (int, error) {
	var id int
	var query string
	n := rand.Intn(2000)
	fmt.Println(n)
	if n < 1000 {
		query = fmt.Sprintf("INSERT INTO %s(user_id, sum, currensy, date_create, date_last_update, status) values($1, $2, $3, NOW(), NOW(), 'Новый') RETURNING id", transactionsTable)
	} else {
		query = fmt.Sprintf("INSERT INTO %s(user_id, sum, currensy, date_create, date_last_update, status) values($1, $2, $3, NOW(), NOW(), 'Ошибка') RETURNING id", transactionsTable)
	}
	row := tran.db.QueryRow(query, transaction.User_id, transaction.Sum, transaction.Currensy)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (tran *transactionItemPostgres) GetTransaction(id string) (test.Transactions, error) {
	var t test.Transactions

	query := fmt.Sprintf("SELECT * FROM %s WHERE ID = $1", transactionsTable)
	err := tran.db.Get(&t, query, id)

	return t, err
}

func (tran *transactionItemPostgres) UpdateTransaction(id string, discription string) error {
	var t test.Transactions
	query := fmt.Sprintf("SELECT * FROM %s WHERE ID = $1", transactionsTable)

	err := tran.db.Get(&t, query, id)

	if err != nil {
		return err
	}

	if t.Status == "Новый" {
		query = fmt.Sprintf("UPDATE %s SET status = $1, date_last_update = NOW() WHERE ID = $2", transactionsTable)
	} else {
		err = errors.New("nе возможно изменить статус транзакции")
		return err
	}
	_, err = tran.db.Exec(query, discription, id)

	return err
}
