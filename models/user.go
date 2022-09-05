package models

import "time"

type User struct {
	ID        int       `json:"id" gorm:"PRIMARY_KEY"`
	Fullname  string    `json:"fullname" gorm:"type: varchar(255)"`
	Email     string    `json:"email" gorm:"type: varchar(255)"`
	Password  string    `json:"password" gorm:"type: varchar(255)"`
	Status    string    `json:"status" gorm:"default:customer"`
	Address   string    `json:"address" gorm:"type: varchar(255)"`
	Phone     string    `json:"phone" gorm:"type: varchar(255)"`
	City      string    `json:"city" gorm:"type: varchar(255)"`
	Image     string    `json:"image" gorm:"default:https://res.cloudinary.com/dqwv0exem/image/upload/v1662340465/waysbeans/default_wr7jvq.jpg"`
	PostCode  string    `json:"postcode" gorm:"type: varchar(255)"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}
