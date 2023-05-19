package spec

type ProductResponse map[int32]Product

type ProductRequest struct {
	ID       int32 `json:"id"`
	Quantity int32 `json:"quantity"`
}
