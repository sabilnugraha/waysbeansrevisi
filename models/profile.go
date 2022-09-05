package models

import "time"

type Profile struct {
	ID        int       `json:"id" gorm:"primary_key:auto_increment"`
	Address   string    `json:"address" gorm:"type: varchar(255)"`
	Phone     int       `json:"phone" gorm:"type: varchar(255)"`
	Image     string    `json:"image" gorm:"type: varchar(255)"`
	UserID    int       `json:"user_id"`
	PostCode  string    `json:"post" gorm:"type: varchar(255)"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

type ProfileResponse struct {
	Address  string `json:"address"`
	Phone    string `json:"phone"`
	Image    string `json:"image" `
	UserID   int    `json:"user_id"`
	PostCode string `json:"post"`
}

func (ProfileResponse) TableName() string {
	return "profiles"
}
