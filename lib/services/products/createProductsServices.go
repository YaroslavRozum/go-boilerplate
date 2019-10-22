package products

import (
	"github.com/YaroslavRozum/go-boilerplate/lib/models"
	"github.com/YaroslavRozum/go-boilerplate/lib/runner"
	"gopkg.in/go-playground/validator.v9"
)

type Services struct {
	List func() runner.Service
}

func CreateProductsServices(mappers models.Mappers, validate *validator.Validate) Services {
	return Services{
		List: func() runner.Service {
			return &ProductsList{mappers, validate}
		},
	}
}
