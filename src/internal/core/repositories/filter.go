package repositories

type Filter struct {
	Key      string
	Value    string
	Partial  bool
	IsString bool
}

func NewFilter(key, value string, isString, partial bool) *Filter {
	return &Filter{
		Key:      key,
		Value:    value,
		IsString: isString,
		Partial:  partial,
	}
}
