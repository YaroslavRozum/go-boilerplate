package controllers

import (
	"net/http"
	"strconv"

	"github.com/YaroslavRozum/go-boilerplate/services/users"
)

type UsersControllers struct {
	List http.HandlerFunc
}

func CreateUsersControllers() UsersControllers {
	return UsersControllers{
		List: NewController(
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
	}
}
