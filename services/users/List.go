package users

import (
	"github.com/YaroslavRozum/go-boilerplate/errors"
	"github.com/YaroslavRozum/go-boilerplate/models"
	"github.com/YaroslavRozum/go-boilerplate/services"
	"github.com/YaroslavRozum/go-boilerplate/services/sessions"
	"github.com/YaroslavRozum/go-boilerplate/services/utils"
)

var validate = services.Validate

type UsersListRequest struct {
	Search string
	Offset int `validate:"omitempty,gte=0"`
	Limit  int `validate:"omitempty,gte=0"`
}

type UsersListResponse struct {
	Users []*models.User `json:"users"`
}

type UsersList struct {
	context *sessions.Context
}

func (uL *UsersList) SetContext(data interface{}) {
	ctx := data.(*sessions.Context)
	uL.context = ctx
}

func (uL *UsersList) Execute(data interface{}) (interface{}, error) {
	payload := data.(*UsersListRequest)
	offset := uint64(payload.Offset)
	limit := uint64(payload.Limit)
	mapper := models.DefaultUserMapper

	users, err := mapper.FindAll(nil, limit, offset)
	if err != nil {
		return nil, &errors.Error{Status: 0, Reason: err.Error()}
	}

	responseData := []*models.User{}

	for _, user := range users {
		userToAppend := utils.DumpUser(user)
		responseData = append(responseData, &userToAppend)
	}

	return UsersListResponse{responseData}, nil
}

func (uL *UsersList) Validate(data interface{}) error {
	err := validate.Struct(data)
	if err != nil {
		return &errors.Error{Status: 0, Reason: err.Error()}
	}
	return nil
}
