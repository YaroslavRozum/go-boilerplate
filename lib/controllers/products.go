package controllers

import (
	"net/http"
	"strconv"

	"github.com/YaroslavRozum/go-boilerplate/lib/services/products"
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
				offset, _ := strconv.ParseUint(query.Get("offset"), 10, 64)
				limit, _ := strconv.ParseUint(query.Get("limit"), 10, 64)
				requestData := products.ProductsListRequest{
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
