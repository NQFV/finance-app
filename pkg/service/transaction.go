package service

import (
	"github.com/NQFV/p/pkg/models"
	"github.com/NQFV/p/pkg/repository"
)

type TransactionService struct {
	repo repository.Transaction
}

func NewTransactionService(repo repository.Transaction) *TransactionService {
	return &TransactionService{repo: repo}
}

func (s *TransactionService) Create(userId int, list models.Transaction) (int, error) {
	return s.repo.Create(userId, list)
}

func (s *TransactionService) GetAll(userId int) ([]models.Transaction, error) {
	return s.repo.GetAll(userId)
}

func (s *TransactionService) GetById(userId, transactionId int) (models.Transaction, error) {
	return s.repo.GetById(userId, transactionId)
}

func (s *TransactionService) Update(userId, transactionId int, input models.Transaction) error {
	return s.repo.Update(userId, transactionId, input)
}

func (s *TransactionService) Delete(userId, transactionId int) error {
	return s.repo.Delete(userId, transactionId)
}
