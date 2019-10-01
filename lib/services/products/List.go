package products

import (
	"github.com/YaroslavRozum/go-boilerplate/lib/errors"
	"github.com/YaroslavRozum/go-boilerplate/lib/models"
	"github.com/YaroslavRozum/go-boilerplate/lib/services"
	"github.com/YaroslavRozum/go-boilerplate/lib/services/utils"
)

var validate = services.Validate

type ProductsListRequest struct {
	Search string
	Offset uint64 `validate:"gte=0,omitempty"`
	Limit  uint64 `validate:"gte=0,omitempty,max=100"`
}

type ProductsListResponse struct {
	Products []*models.Product `json:"products"`
}

type ProductsList struct{}

func (pL *ProductsList) Execute(data interface{}) (interface{}, error) {
	payload := data.(*ProductsListRequest)
	offset := payload.Offset
	limit := payload.Limit
	productMapper := models.DefaultProductMapper

	products, err := productMapper.FindAll(nil, limit, offset)
	if err != nil {
		return nil, &errors.Error{Status: 0, Reason: err.Error()}
	}

	responseData := []*models.Product{}

	for _, product := range products {
		productToAppend := utils.DumpProduct(product)
		responseData = append(responseData, &productToAppend)
	}

	return ProductsListResponse{responseData}, nil
}

func (pL *ProductsList) Validate(data interface{}) error {
	err := validate.Struct(data)
	if err != nil {
		return &errors.Error{Status: 0, Reason: err.Error()}
	}
	return nil
}
