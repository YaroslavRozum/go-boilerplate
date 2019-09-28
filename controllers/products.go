package controllers

import (
	"net/http"
	"strconv"

	"github.com/YaroslavRozum/go-boilerplate/services/products"
)

type ProductsControllers struct {
	List http.HandlerFunc
}

func CreateProductsControllers() ProductsControllers {
	return ProductsControllers{
		List: NewController(
			NewServiceRunnerCreator(&products.ProductsList{}),
			func(r *http.Request) (interface{}, error) {
				query := r.URL.Query()
				offset, _ := strconv.Atoi(query.Get("offset"))
				limit, _ := strconv.Atoi(query.Get("limit"))
				requestData := &products.ProductsListRequest{
					Search: query.Get("search"),
					Offset: offset,
					Limit:  limit,
				}
				return requestData, nil
			},
			defaultJsonResponse,
		),
	}
}
