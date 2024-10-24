package domain

import "time"

type Sale struct {
	Id            int        `json:"id"`
	TotalPrice    float32    `json:"total_price"`
	Discount      *float32   `json:"discount"`
	Comments      *string    `json:"comments"`
	TypeOfPayment int        `json:"type_of_payment"`
	CreatedAt     time.Time  `json:"created_at"`
	Employee      *Employee  `json:"employee_id"`
	Customer      *Customer  `json:"customer_id"`
	Items         []SaleItem `json:"items"`
}

type SaleItem struct {
	Id       int      `json:"id"`
	Quantity int      `json:"quantity"`
	Product  *Product `json:"product_id"`
	Sale     *Sale    `json:"sale_id"`
}
