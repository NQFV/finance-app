package service

import (
	"time"

	"github.com/NQFV/p/pkg/models"
	"github.com/NQFV/p/pkg/repository"
)

type TransactionService struct {
	repo repository.Transaction
}

func NewTransactionService(repo repository.Transaction) *TransactionService {
	return &TransactionService{repo: repo}
}

func (s *TransactionService) Create(userId int, transaction models.Transaction) (int, error) {
	if transaction.Date.IsZero() {
		transaction.Date = time.Now().UTC()
	}
	return s.repo.Create(userId, transaction)
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
