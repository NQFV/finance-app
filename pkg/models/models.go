package models

import "time"

type Transaction struct {
	Id       int       `json:"id" db:"transaction_id"`
	Type     string    `json:"type" db:"type"`
	Date     time.Time `json:"date" db:"date"`
	Amount   int       `json:"amount" db:"amount"`
	Category Category  `json:"category" db:"category"`
	UserId   int       `json:"userId" db:"user_id"`
}

type Category struct {
	Id   int    `json:"id" db:"category_id"`
	Name string `json:"name" db:"name"`
}
