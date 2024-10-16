package responses

type PaginatedResponse[T any] struct {
	Items      []T `json:"items"`
	Page       int `json:"page"`
	PageSize   int `json:"page_size"`
	TotalCount int `json:"total_count"`
}

func NewPaginatedResponse[T any](items []T, page int, pageSize int, totalCount int) *PaginatedResponse[T] {
	return &PaginatedResponse[T]{
		Items:      items,
		Page:       page,
		PageSize:   pageSize,
		TotalCount: totalCount,
	}
}
