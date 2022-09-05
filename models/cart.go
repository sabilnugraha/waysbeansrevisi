package models

import "time"

type Cart struct {
	ID            int       `json:"id" gorm:"PRIMARY_KEY"`
	Product_ID    int       `json:"product_id"`
	TransactionID *int      `json:"transaction_id"`
	Product       Product   `json:"product" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	UserID        int       `json:"user_id"`
	User          User      `json:"user"`
	SubTotal      *int      `json:"subtotal"`
	Qty           *int      `json:"qty" gorm:"default:1"`
	CreatedAt     time.Time `json:"-"`
	UpdatedAt     time.Time `json:"updated_at"`
	Stockproduct  *int      `json:"stockproduct"`
}

type CartResponse struct {
	ID            int                `json:"id"`
	Total         int                `json:"total"`
	TransactionID int                `json:"transaction_id"`
	ProductID     int                `json:"product_id"`
	Product       ProductTransaction `json:"product"`
	Qty           int                `json:"qty"`
	UserID        int                `json:"user_id"`
}
