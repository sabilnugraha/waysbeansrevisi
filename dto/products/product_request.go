package productsdto

type CreateProductRequest struct {
	Name  string `json:"name" form:"title" gorm:"type: varchar(255)"`
	Price int    `json:"price" form:"price" gorm:"type: int"`
	Stock int    `json:"stock" gorm:"type: int"`
	Desc  string `json:"desc" form:"desc" gorm:"type: varchar(255)"`
}

type UpdateProductRequest struct {
	Name  string `json:"name" form:"title" gorm:"type: varchar(255)"`
	Price int    `json:"price" form:"price" gorm:"type: int"`
	Stock int    `json:"stock" gorm:"type: int"`
	Desc  string `json:"desc" form:"desc" gorm:"type: varchar(255)"`
	Image string `json:"image" gorm:"type: varchar(255)"`
}

type UpdateStockRequest struct {
	Stock int `json:"stock" gorm:"type: int"`
}
