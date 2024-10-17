package domain

import (
	"time"
)

type Product struct {
	Id            int        `json:"id"`
	Name          string     `json:"name"`
	PurchasePrice float32    `json:"purchase_price"`
	SalePrice     float32    `json:"sale_price"`
	Stock         int        `json:"stock"`
	IsFeedstock   bool       `json:"is_feedstock"`
	Barcode       string     `json:"barcode"`
	CreatedAt     time.Time  `json:"created_at"`
	CreatedBy     *Employee  `json:"created_by"`
	Category      *Category  `json:"category"`
	UpdatedAt     *time.Time `json:"updated_at"`
	UpdatedBy     *Employee  `json:"updated_by"`
}
