package controllers

import (
	"net/http"
	"strconv"

	"github.com/YaroslavRozum/go-boilerplate/lib/services/users"
)

type UsersControllers struct {
	List   http.HandlerFunc
	Create http.HandlerFunc
}

func CreateUsersControllers() UsersControllers {
	return UsersControllers{
		List: NewController(
			NewServiceRunnerCreator(&users.UsersList{}),
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
		Create: NewController(
			NewServiceRunnerCreator(&users.UsersCreate{}),
			defaultPayloadBuilder(users.UsersCreateRequest{}),
			defaultJsonResponse,
		),
	}
}
