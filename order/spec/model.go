package spec

type Order struct {
	ID             int32   `json:"id"`
	Quantity       int32   `json:"quantity"`
	Status         string  `json:"order_status"`
	Amount         float32 `json:"order_amount"`
	Discount       float32 `json:"discount"`
	FinalAmount    float32 `json:"final_amount"`
	DispatchedDate string  `json:"dispatched_date"`
}

type OrderItem struct {
	ID       int32 `json:"product_id"`
	Quantity int32 `json:"quantity"`
}

var Orders map[int32]Order
