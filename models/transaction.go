package models

import "time"

type Transaction struct {
	ID        int       `json:"id" gorm:"PRIMARY_KEY"`
	UserID    int       `json:"user_id"`
	User      User      `json:"user"`
	Carts     []Cart    `json:"carts"`
	Status    string    `json:"status"`
	Total     int       `json:"total"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"updated_at"`
}

type TransactionResponse struct {
	UserID int `json:"-"`
	User   int `json:"user_id"`
}

func (TransactionResponse) TableName() string {
	return "profiles"
}
