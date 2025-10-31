package repository

import (
	"github.com/NQFV/p/pkg/models"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user models.User) (int, error)
	GetUser(username, password string) (models.User, error)
}

type Transaction interface {
	Create(userId int, list models.Transaction) (int, error)
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

type Repository struct {
	Authorization
	Transaction
	Category
	Anal
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		Transaction:   NewTransactionPostgres(db),
		Category:      NewCategoryPostgres(db),
	}
}
