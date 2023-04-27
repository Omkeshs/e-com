package spec

type OrderRequest struct {
	Items []OrderItem `json:"items"`
}

type UpdateOrderRequest struct {
	ID     int32  `json:"order_id"`
	Date   string `json:"order_date"`
	Status string `json:"order_status"`
}

type UpdateProductRequest struct {
	ID       int32 `json:"id"`
	Quantity int32 `json:"quantity"`
}
