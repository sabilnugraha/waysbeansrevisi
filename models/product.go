package models

import "time"

type Product struct {
	ID        int       `json:"id" gorm:"PRIMARY_KEY"`
	Name      string    `json:"name" form:"name" gorm:"type: varchar(255)"`
	Price     int       `json:"price" gorm:"type: int"`
	Stock     int       `json:"stock" gorm:"type: int"`
	Image     string    `json:"image"  form:"image" gorm:"type: varchar(255)"`
	Desc      string    `json:"desc" form:"desc" gorm:"type: varchar(255)"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

type ProductResponse struct {
	Name  string `json:"name"`
	Price int    `json:"price"`
	Image string `json:"image"`
	Stock int    `json:"stock"`
	Desc  string `json:"desc"`
}

type ProductTransaction struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
	Image string `json:"image"`
}

func (ProductResponse) TableName() string {
	return "products"
}
