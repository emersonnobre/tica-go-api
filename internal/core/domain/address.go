package domain

type Address struct {
	Id           int     `json:"id"`
	Street       string  `json:"street"`
	Neighborhood string  `json:"neighborhood"`
	Cep          *string `json:"cep"`
	CustomerId   int     `json:"customer_id"`
}
