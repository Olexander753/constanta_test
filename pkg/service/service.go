package service

import (
	test "github.com/Olexander753/constanta_test"
	"github.com/Olexander753/constanta_test/pkg/repository"
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

type Service struct {
	TransactionsList
	TransactionItem
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		TransactionItem:  NewTransactionService(repo.TransactionItem),
		TransactionsList: NewTransactionListService(repo.TransactionsList),
	}
}
