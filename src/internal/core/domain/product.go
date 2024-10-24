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
	Barcode       string     `json:"barcode"`
	Category      *Category  `json:"category"`
	Active        bool       `json:"-"`
	CreatedAt     time.Time  `json:"created_at"`
	CreatedBy     *Employee  `json:"created_by"`
	UpdatedAt     *time.Time `json:"updated_at"`
	UpdatedBy     *Employee  `json:"updated_by"`
	IsFeedstock   bool       `json:"is_feedstock"`
}

func NewEmptyProduct() *Product {
	var product Product
	product.Category = &Category{}
	product.CreatedBy = &Employee{}
	product.UpdatedBy = &Employee{}
	return &product
}

func (p *Product) Update(stock int, purchasePrice float32, updatedBy *Employee) {
	p.PurchasePrice = purchasePrice
	p.Stock = stock
	now := time.Now()
	p.UpdatedAt = &now
	p.UpdatedBy = updatedBy
}
