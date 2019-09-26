package users

import (
	"github.com/YaroslavRozum/go-boilerplate/errors"
	"github.com/YaroslavRozum/go-boilerplate/services"
	"github.com/YaroslavRozum/go-boilerplate/services/sessions"
)

var validate = services.Validate

type UsersListRequest struct {
	Search string
	Offset int `validate:"gte=0"`
	Limit  int `validate:"required,gte=0"`
}

type UsersList struct {
	context *sessions.Context
}

func (uL *UsersList) SetContext(ctx interface{}) {
	uL.context = ctx.(*sessions.Context)
}

func (uL *UsersList) Execute(data interface{}) interface{} {
	payload := data.(*UsersListRequest)
	offset := payload.Offset
	limit := payload.Limit
	lastIndex := offset + limit
	if offset+limit > 20 {
		lastIndex = 20
	}
	return []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}[payload.Offset:lastIndex]
}

func (uL *UsersList) Validate(data interface{}) error {
	err := validate.Struct(data)
	if err != nil {
		return &errors.Error{Status: 0, Reason: err.Error()}
	}
	return nil
}
