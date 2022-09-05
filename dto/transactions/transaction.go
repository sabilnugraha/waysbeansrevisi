package transactiondto

type CreateTransaction struct {
	Total int `json:"total"`
}

type UpdateTransaction struct {
	UserID int    `json:"user_id" form:"user_id"`
	Status string `json:"status"`
	Total  int    `json:"total"`
}

type TransactionResponse struct {
	ID     int    `json:"id"`
	UserID int    `json:"user_id" form:"user_id"`
	Status string `json:"status"`
	Total  int    `json:"total"`
}
