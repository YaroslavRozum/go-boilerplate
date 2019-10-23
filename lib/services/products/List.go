package products

import (
	"github.com/YaroslavRozum/go-boilerplate/lib/errors"
	"github.com/YaroslavRozum/go-boilerplate/lib/models"
	"github.com/YaroslavRozum/go-boilerplate/lib/services/utils"
	"gopkg.in/go-playground/validator.v9"
)

type ProductsListRequest struct {
	Search string
	Offset uint64 `validate:"gte=0,omitempty"`
	Limit  uint64 `validate:"gte=0,omitempty,max=100"`
}

type ProductsListResponse struct {
	Products []models.Product `json:"products"`
}

type ProductsList struct {
	mappers  models.Mappers
	validate *validator.Validate
}

func (pL *ProductsList) Execute(data interface{}) (interface{}, error) {
	payload := data.(ProductsListRequest)
	offset := payload.Offset
	limit := payload.Limit
	productMapper := pL.mappers.ProductMapper

	products, err := productMapper.FindAll(nil, limit, offset)
	if err != nil {
		return nil, &errors.Error{Status: 0, Reason: err.Error()}
	}

	responseData := make([]models.Product, 0, len(products))

	for _, product := range products {
		productToAppend := utils.DumpProduct(product)
		responseData = append(responseData, productToAppend)
	}

	return ProductsListResponse{responseData}, nil
}

func (pL *ProductsList) Validate(data interface{}) error {
	err := pL.validate.Struct(data)
	if err != nil {
		return &errors.Error{Status: 0, Reason: err.Error()}
	}
	return nil
}
