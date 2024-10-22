package domain

import "time"

type Transaction struct {
	Id        int       `json:"id"`
	Reason    string    `json:"reason"`
	Quantity  int       `json:"quantity"`
	Type      int       `json:"type"`
	CreatedAt time.Time `json:"created_at"`
	CreatedBy *Employee `json:"created_by"`
	Product   *Product  `json:"product"`
}

func NewTransaction(reason string, quantity, typee int, createdAt time.Time, createdBy *Employee, product *Product) *Transaction {
	return &Transaction{
		Id:        0,
		Reason:    reason,
		Quantity:  quantity,
		Type:      typee,
		CreatedAt: createdAt,
		Product:   product,
	}
}
