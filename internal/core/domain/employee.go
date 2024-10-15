package domain

type Employee struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	CreatedAt string `json:"-"`
}
