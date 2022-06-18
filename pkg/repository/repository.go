package repository

import (
	test "github.com/Olexander753/constanta_test"
	"github.com/jmoiron/sqlx"
)

type TransactionsList interface {
	GetListByID(user_id string) ([]test.Transactions, error)
	GetListByEmail(email string) ([]test.Transactions, error)
}

type TransactionItem interface {
	CreateTransaction(tran test.Transaction) (int, error)
	GetTransaction(id string) (test.Transactions, error)
	UpdateTransaction(id string, discription string) error
}

type Repository struct {
	TransactionsList
	TransactionItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		TransactionItem:  NewTransactionItemPostgres(db),
		TransactionsList: NewTransactionsListPostgres(db),
	}
}
