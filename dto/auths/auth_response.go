package authdto

type RegisterResponse struct {
	Fullname string `gorm:"type: varchar(255)" json:"fullname"`
	Email    string `gorm:"type: varchar(255)" json:"email"`
}

type LoginResponse struct {
	ID       int    `json:"id"`
	Fullname string `gorm:"type: varchar(255)" json:"fullname"`
	Email    string `gorm:"type: varchar(255)" json:"email"`
	Token    string `gorm:"type: varchar(255)" json:"token"`
	Status   string `json:"status"`
	Address  string `gorm:"type: varchar(255)"`
	Phone    string `gorm:"type: varchar(255)"`
	City     string `gorm:"type: varchar(255)"`
	PostCode string `gorm:"type: varchar(255)"`
	Image    string `json:"image"`
}

type CheckAuthResponse struct {
	ID       int    `json:"id"`
	Fullname string `gorm:"type: varchar(255)" json:"fullname"`
	Email    string `gorm:"type: varchar(255)" json:"email"`
	Status   string `gorm:"type: varchar(50)"  json:"status"`
	Address  string `gorm:"type: varchar(255)"`
	Phone    string `gorm:"type: varchar(255)"`
	City     string `gorm:"type: varchar(255)"`
	PostCode string `gorm:"type: varchar(255)"`
	Image    string `json:"image"`
}

type UpdateUserResponse struct {
	Address  string `gorm:"type: varchar(255)"`
	Phone    string `gorm:"type: varchar(255)"`
	City     string `gorm:"type: varchar(255)"`
	PostCode string `gorm:"type: varchar(255)"`
}
