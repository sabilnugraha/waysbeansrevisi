package cartdto

type CreateCart struct {
	ID         int    `json:"id"`
	UserID     int    `json:"user_id"`
	Product_ID int    `json:"product_id"`
	SubTotal   *int   `json:"subtotal"`
	Status     string `json:"status"`
}

type UpdateQtyRequest struct {
	Qty          int `json:"qty"`
	SubTotal     int `json:"subtotal"`
	Stockproduct int `json:"stockproduct"`
}

type UpdateCartRequest struct {
	TransactionID int `json:"transaction_id"`
}
