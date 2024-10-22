package requests

type ProductOutflow struct {
	ProductId int    `json:"product_id"`
	Quantity  int    `json:"quantity"`
	Reason    string `json:"reason"`
	CreatedBy int    `json:"created_by"`
}
