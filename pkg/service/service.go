package service

import (
	"github.com/NQFV/p/pkg/models"
	"github.com/NQFV/p/pkg/repository"
)

type Authorization interface {
	CreateUser(user models.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type Transaction interface {
	Create(userId int, transaction models.Transaction) (int, error)
	GetAll(userId int) ([]models.Transaction, error)
	GetById(userId, transactionId int) (models.Transaction, error)
	Update(userId, transactionId int, input models.Transaction) error
	Delete(userId, transactionId int) error
}

type Category interface {
	Create(Id int, category models.Category) (int, error)
	GetAll(userId int) ([]models.Category, error)
	Update(userId, categoryId int, input models.Category) error
	Delete(userId, categoryId int) error
}

type Anal interface {
}

type Service struct {
	Authorization
	Transaction
	Category
	Anal
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		Transaction:   NewTransactionService(repos.Transaction),
		Category:      NewCategoryService(repos.Category),
	}

}
