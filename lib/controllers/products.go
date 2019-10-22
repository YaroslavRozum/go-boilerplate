package controllers

import (
	"net/http"
	"strconv"

	"github.com/YaroslavRozum/go-boilerplate/lib/runner"
	"github.com/YaroslavRozum/go-boilerplate/lib/services/products"
)

type ProductsControllers struct {
	List http.HandlerFunc
}

func CreateProductsControllers(p products.Services) ProductsControllers {
	return ProductsControllers{
		List: runner.NewController(
			runner.NewServiceRunnerCreator(p.List),
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
