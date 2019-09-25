package products

import (
	"github.com/YaroslavRozum/go-boilerplate/errors"
	"github.com/YaroslavRozum/go-boilerplate/services"
)

var validate = services.Validate

type ProductsListRequest struct {
	Search string
	Offset int `validate:"gte=0"`
	Limit  int `validate:"required,gte=0"`
}

type ProductsList struct{}

func (pL *ProductsList) Execute(data interface{}) interface{} {
	payload := data.(*ProductsListRequest)
	offset := payload.Offset
	limit := payload.Limit
	lastIndex := offset + limit
	if offset+limit > 3 {
		lastIndex = 3
	}

	return []string{"str", "lll", "222"}[payload.Offset:lastIndex]
}

func (pL *ProductsList) Validate(data interface{}) error {
	err := validate.Struct(data)
	if err != nil {
		return &errors.Error{Status: 0, Reason: err.Error()}
	}
	return nil
}
