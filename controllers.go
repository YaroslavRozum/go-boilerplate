package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/YaroslavRozum/go-boilerplate/services/products"
	"github.com/YaroslavRozum/go-boilerplate/services/users"
)

type Controllers struct {
	UsersList    http.HandlerFunc
	ProductsList http.HandlerFunc
}

func defaultJsonResponse(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	encoder.Encode(data)
}

func CreateControllers() *Controllers {
	return &Controllers{
		UsersList: NewController(
			NewServiceRunnerCreator(&users.UsersList{}),
			func(r *http.Request) interface{} {
				query := r.URL.Query()
				offset, _ := strconv.Atoi(query.Get("offset"))
				limit, _ := strconv.Atoi(query.Get("limit"))
				return &users.UsersListRequest{
					Search: query.Get("search"),
					Offset: offset,
					Limit:  limit,
				}
			},
			defaultJsonResponse,
		),
		ProductsList: NewController(
			NewServiceRunnerCreator(&products.ProductsList{}),
			func(r *http.Request) interface{} {
				query := r.URL.Query()
				offset, _ := strconv.Atoi(query.Get("offset"))
				limit, _ := strconv.Atoi(query.Get("limit"))
				return &products.ProductsListRequest{
					Search: query.Get("search"),
					Offset: offset,
					Limit:  limit,
				}
			},
			defaultJsonResponse,
		),
	}
}
