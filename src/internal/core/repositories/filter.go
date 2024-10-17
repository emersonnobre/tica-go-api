package repositories

type Filter struct {
	Key     string
	Value   string
	Partial bool
	Type    string
}

func NewFilter(key, value, typee string, partial bool) *Filter {
	return &Filter{
		Key:     key,
		Value:   value,
		Type:    typee,
		Partial: partial,
	}
}
