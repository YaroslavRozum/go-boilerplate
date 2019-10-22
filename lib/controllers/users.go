package controllers

import (
	"net/http"
	"strconv"

	"github.com/YaroslavRozum/go-boilerplate/lib/runner"
	"github.com/YaroslavRozum/go-boilerplate/lib/services/users"
)

type UsersControllers struct {
	List   http.HandlerFunc
	Create http.HandlerFunc
}

func CreateUsersControllers(u users.Services) UsersControllers {
	return UsersControllers{
		List: runner.NewController(
			runner.NewServiceRunnerCreator(u.List),
			func(r *http.Request) (interface{}, error) {
				query := r.URL.Query()
				offset, _ := strconv.ParseUint(query.Get("offset"), 10, 64)
				limit, _ := strconv.ParseUint(query.Get("limit"), 10, 64)
				requestData := users.UsersListRequest{
					Search: query.Get("search"),
					Offset: offset,
					Limit:  limit,
				}
				return requestData, nil
			},
			defaultJsonResponse,
		),
		Create: runner.NewController(
			runner.NewServiceRunnerCreator(u.Create),
			defaultPayloadBuilder(users.UsersCreateRequest{}),
			defaultJsonResponse,
		),
	}
}
