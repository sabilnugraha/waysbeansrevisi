package productsdto

type ProductResponse struct {
	ID    int    `json:"id"`
	Name  string `json:"name" gorm:"type: varchar(255)"`
	Price int    `json:"price" gorm:"type: varchar(255)"`
	Image string `json:"image" form:"image" validate:"required"`
	Stock int    `json:"stock" gorm:"type: int"`
	Desc  string `json:"desc" form:"desc" gorm:"type: varchar(255)"`
}
