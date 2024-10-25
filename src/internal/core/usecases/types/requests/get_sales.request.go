package requests

import (
	"errors"
	"strings"
)

type GetSalesRequest struct {
	limit   int
	offset  int
	orderBy string
	order   string
}

func NewGetSalesRequest(limit, offset int, orderBy, order string) (*GetSalesRequest, error) {
	sale := &GetSalesRequest{
		limit:   limit,
		offset:  offset,
		orderBy: orderBy,
		order:   order,
	}

	if err := sale.validateObject(); err != nil {
		return nil, err
	}
	return sale, nil
}

func (r *GetSalesRequest) validateObject() error {
	if r.limit == 0 {
		r.limit = 10
	}
	r.orderBy = strings.TrimSpace(r.orderBy)
	if r.order != "ASC" && r.order != "DESC" {
		return errors.New("ordenação deve ser ASC ou DESC")
	}
	return nil
}

func (r *GetSalesRequest) Limit() int {
	return r.limit
}

func (r *GetSalesRequest) Offset() int {
	return r.offset
}

func (r *GetSalesRequest) OrderBy() string {
	return r.orderBy
}

func (r *GetSalesRequest) Order() string {
	return r.order
}
