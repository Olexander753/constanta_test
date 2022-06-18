package service

import (
	test "github.com/Olexander753/constanta_test"
	"github.com/Olexander753/constanta_test/pkg/repository"
)

type TransactionsListService struct {
	repo repository.TransactionsList
}

func NewTransactionListService(repo repository.TransactionsList) *TransactionsListService {
	return &TransactionsListService{repo: repo}
}

func (s *TransactionsListService) GetListByID(user_id string) ([]test.Transactions, error) {
	return s.repo.GetListByID(user_id)
}

func (s *TransactionsListService) GetListByEmail(email string) ([]test.Transactions, error) {
	return s.repo.GetListByEmail(email)
}
