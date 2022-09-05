package cartdto

type CartResponse struct {
	ID       int `json:"id"`
	SubTotal int `json:"subtotal"`
}
