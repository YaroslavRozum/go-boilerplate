package users

import (
	sq "github.com/Masterminds/squirrel"
	"github.com/YaroslavRozum/go-boilerplate/lib/errors"
	"github.com/YaroslavRozum/go-boilerplate/lib/models"
	"github.com/YaroslavRozum/go-boilerplate/lib/services"
	"github.com/YaroslavRozum/go-boilerplate/lib/services/sessions"
	"github.com/YaroslavRozum/go-boilerplate/lib/services/utils"
)

var validate = services.Validate

type UsersListRequest struct {
	Search string
	Offset uint64 `validate:"omitempty,gte=0"`
	Limit  uint64 `validate:"omitempty,gte=0,max=100"`
}

type UsersListResponse struct {
	Users []*models.User `json:"users"`
}

type UsersList struct {
	context sessions.Context
}

func (uL *UsersList) SetContext(data interface{}) {
	ctx := data.(sessions.Context)
	uL.context = ctx
}

func (uL *UsersList) Execute(data interface{}) (interface{}, error) {
	payload := data.(UsersListRequest)
	offset := payload.Offset
	limit := payload.Limit
	userMapper := models.DefaultUserMapper
	ctx := uL.context
	users, err := userMapper.FindAll(sq.NotEq{"id": ctx.ID}, limit, offset)
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
