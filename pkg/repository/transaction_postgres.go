package repository

import (
	"fmt"
	"time"

	"github.com/NQFV/p/pkg/models"
	"github.com/jmoiron/sqlx"
)

type TransactionPostgres struct {
	db *sqlx.DB
}

func NewTransactionPostgres(db *sqlx.DB) *TransactionPostgres {
	return &TransactionPostgres{db: db}
}

func (r *TransactionPostgres) Create(userId int, transaction models.Transaction) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	newDate := time.Now()
	var id int
	createTransactionQuery := fmt.Sprintf("INSERT INTO %s (type, date, amount, category_id, user_id) VALUES ($1, $2, $3, $4, $5) RETURNING transaction_id", transactionsTable)
	row := tx.QueryRow(createTransactionQuery, transaction.Type, newDate, transaction.Amount, transaction.Category.Id, userId)
	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}

	createUsersTransactionQuery := fmt.Sprintf("INSERT INTO %s (user_id, transaction_id) VALUES ($1, $2)", users_transactionsTable)
	_, err = tx.Exec(createUsersTransactionQuery, userId, id)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return id, tx.Commit()
}

func (r *TransactionPostgres) GetAll(userId int) ([]models.Transaction, error) {
	var transactions []models.Transaction

	query := fmt.Sprintf(`SELECT t.transaction_id, t.type, t.date, t.amount, t.user_id, c.category_id, c.name  FROM %s t INNER JOIN %s ul ON t.transaction_id = ul.transaction_id LEFT JOIN %s c ON t.category_id = c.category_id WHERE ul.user_id = $1`,
		transactionsTable, users_transactionsTable, categoriesTable)
	rows, err := r.db.Query(query, userId)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var trans models.Transaction
		var catId int
		var catName string

		if err := rows.Scan(&trans.Id, &trans.Type, &trans.Date, &trans.Amount, &trans.UserId, &catId, &catName); err != nil {
			return nil, err
		}

		trans.Category = models.Category{Id: catId, Name: catName}
		transactions = append(transactions, trans)
	}

	return transactions, err
}

func (r *TransactionPostgres) GetById(userId, transactionId int) (models.Transaction, error) {
	var transaction models.Transaction
	var catId int
	var catName string

	query := fmt.Sprintf(`SELECT t.transaction_id, t.type, t.date, t.amount, t.user_id, c.category_id, c.name FROM %s t
								INNER JOIN %s ul on t.transaction_id = ul.transaction_id LEFT JOIN %s c ON t.category_id = c.category_id WHERE ul.user_id = $1 AND ul.transaction_id = $2`,
		transactionsTable, users_transactionsTable, categoriesTable)
	err := r.db.QueryRow(query, userId, transactionId).Scan(&transaction.Id, &transaction.Type, &transaction.Date, &transaction.Amount, &transaction.UserId, &catId, &catName)
	if err != nil {
		return models.Transaction{}, err
	}

	transaction.Category = models.Category{
		Id:   catId,
		Name: catName,
	}

	return transaction, err
}

func (r *TransactionPostgres) Update(userId, transactionId int, input models.Transaction) error {
	query := fmt.Sprintf(`UPDATE %s t SET type = $1, amount = $2, category_id = $3 FROM %s ul WHERE t.transaction_id = ul.transaction_id AND ul.user_id = $4 AND ul.transaction_id = $5`,
		transactionsTable, users_transactionsTable)

	_, err := r.db.Exec(query,
		input.Type,
		input.Amount,
		input.Category.Id,
		userId,
		transactionId,
	)

	return err
}

func (r *TransactionPostgres) Delete(userId, transactionId int) error {
	query := fmt.Sprintf("DELETE FROM %s tl USING %s ul WHERE tl.transaction_id = ul.transaction_id AND ul.user_id=$1 AND ul.transaction_id=$2",
		transactionsTable, users_transactionsTable)
	_, err := r.db.Exec(query, userId, transactionId)

	return err
}
