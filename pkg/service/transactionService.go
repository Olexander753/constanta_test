package service

import (
	test "github.com/Olexander753/constanta_test"
	"github.com/Olexander753/constanta_test/pkg/repository"
)

type TransactionService struct {
	repo repository.TransactionItem
}

func NewTransactionService(repo repository.TransactionItem) *TransactionService {
	return &TransactionService{repo: repo}
}

func (s *TransactionService) CreateTransaction(tran test.Transaction) (int, error) {
	return s.repo.CreateTransaction(tran)
}

func (s *TransactionService) GetTransaction(id string) (test.Transactions, error) {
	return s.repo.GetTransaction(id)
}

func (s *TransactionService) UpdateTransaction(id string, discription string) error {
	return s.repo.UpdateTransaction(id, discription)
}
