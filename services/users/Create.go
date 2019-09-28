package users

import (
	"github.com/YaroslavRozum/go-boilerplate/errors"
	"github.com/YaroslavRozum/go-boilerplate/models"
)

type UsersCreateRequest struct {
	UserName        string `json:"userName" validate:"min=6, max=20"`
	Name            string `json:"name" validate:"min=6, max-20"`
	Surname         string `json:"surname" validate:"min=6, max=20"`
	Email           string `json:"email" validate:"required,email"`
	Password        string `json:"password" validate:"required,min=6,max=20"`
	ConfirmPassword string `json:"confirmPassword" validate:"eqfield=Password"`
}

type UsersCreate struct{}

func (uC *UsersCreate) Execute(data interface{}) (interface{}, error) {
	payload := data.(*UsersCreateRequest)

	mapper := models.DefaultUserMapper
	user, err := mapper.NewUser(
		payload.UserName,
		payload.Name,
		payload.Email,
		payload.Surname,
		payload.Password,
	)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (uC *UsersCreate) Validate(data interface{}) error {
	err := validate.Struct(data)
	if err != nil {
		return &errors.Error{Status: 0, Reason: err.Error()}
	}
	return nil
}
