package repository

import (
	"fmt"

	"github.com/NQFV/p/pkg/models"
	"github.com/jmoiron/sqlx"
)

type CategoryPostgres struct {
	db *sqlx.DB
}

func NewCategoryPostgres(db *sqlx.DB) *CategoryPostgres {
	return &CategoryPostgres{db: db}
}

func (r *CategoryPostgres) Create(Id int, category models.Category) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var id int
	createCategoryQuery := fmt.Sprintf("INSERT INTO %s (name) VALUES ($1) RETURNING category_id", categoriesTable)
	row := tx.QueryRow(createCategoryQuery, category.Name)
	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}

	return id, tx.Commit()
}

func (r *CategoryPostgres) GetAll(userId int) ([]models.Category, error) {
	var categories []models.Category
	query := fmt.Sprintf("SELECT category_id, name FROM %s", categoriesTable)
	err := r.db.Select(&categories, query)
	return categories, err
}

func (r *CategoryPostgres) Update(userId, categoryId int, input models.Category) error {
	query := fmt.Sprintf("UPDATE %s SET name = $1 WHERE category_id = $2", categoriesTable)
	_, err := r.db.Exec(query, input.Name, categoryId)
	return err
}

func (r *CategoryPostgres) Delete(userId, categoryId int) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	deleteLinksQuery := fmt.Sprintf("DELETE FROM %s WHERE category_id = $1", transaction_categoryTable)
	_, err = tx.Exec(deleteLinksQuery, categoryId)
	if err != nil {
		tx.Rollback()
		return err
	}

	updateTransactionsQuery := fmt.Sprintf("UPDATE %s SET category_id = NULL WHERE category_id = $1", transactionsTable)
	_, err = tx.Exec(updateTransactionsQuery, categoryId)
	if err != nil {
		tx.Rollback()
		return err
	}

	deleteCategoryQuery := fmt.Sprintf("DELETE FROM %s WHERE category_id = $1", categoriesTable)
	_, err = tx.Exec(deleteCategoryQuery, categoryId)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}
