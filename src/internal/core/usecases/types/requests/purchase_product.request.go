package requests

type PurchaseProductRequest struct {
	ProductId     int     `json:"product_id"`
	Quantity      int     `json:"quantity"`
	PurchasePrice float32 `json:"purchase_price"`
	CreatedBy     int     `json:"created_by"`
}
