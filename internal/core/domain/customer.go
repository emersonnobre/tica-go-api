package domain

type Customer struct {
	Id        int     `json:"id"`
	Name      string  `json:"name"`
	Phone     *string `json:"phone"`
	Cpf       *string `json:"cpf"`
	Email     *string `json:"email"`
	Instagram *string `json:"instagram"`
	Birthday  *string `json:"birthday"`
	CreatedAt string  `json:"created_at"`
	Active    bool    `json:"-"`
}