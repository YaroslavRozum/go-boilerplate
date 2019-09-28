package controllers

import (
	"net/http"
	"strconv"

	"github.com/YaroslavRozum/go-boilerplate/services/users"
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
				offset, _ := strconv.Atoi(query.Get("offset"))
				limit, _ := strconv.Atoi(query.Get("limit"))
				requestData := &users.UsersListRequest{
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
			defaultPayloadBuilder(&users.UsersCreateRequest{}),
			defaultJsonResponse,
		),
	}
}
