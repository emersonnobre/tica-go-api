package domain

type Employee struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Cpf       string `json:"cpf"`
	CreatedAt string `json:"created_at"`
}
