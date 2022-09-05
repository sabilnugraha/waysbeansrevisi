package authdto

type RegisterRequest struct {
	Fullname string `gorm:"type: varchar(255)" json:"fullname" validate:"required"`
	Email    string `gorm:"type: varchar(255)" json:"email" validate:"required"`
	Password string `gorm:"type: varchar(255)" json:"password"`
}

type LoginRequest struct {
	Email    string `gorm:"type: varchar(255)" json:"email" validate:"required"`
	Password string `gorm:"type: varchar(255)" json:"password" validate:"required"`
}

type UpdateUserRequest struct {
	Address  string `json:"address" gorm:"type: varchar(255)"`
	Phone    string `json:"phone" gorm:"type: varchar(255)"`
	City     string `json:"city" gorm:"type: varchar(255)"`
	PostCode string `json:"postcode" gorm:"type: varchar(255)"`
}

type UpdateProfile struct {
	Image string `json:"image"`
}
