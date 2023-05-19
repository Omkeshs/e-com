package spec

type Product struct {
	ID       int32   `json:"id"`
	Name     string  `json:"product_name"`
	Price    float32 `json:"price"`
	Category int     `json:"category"`
	Quantity int32   `json:"quantity"`
}
