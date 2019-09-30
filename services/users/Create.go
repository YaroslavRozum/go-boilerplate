package users

import (
	sq "github.com/Masterminds/squirrel"
	"github.com/YaroslavRozum/go-boilerplate/errors"
	"github.com/YaroslavRozum/go-boilerplate/models"
	"github.com/YaroslavRozum/go-boilerplate/services"
	"github.com/YaroslavRozum/go-boilerplate/services/utils"
)

type UsersCreateRequest struct {
	UserName        string `json:"username" validate:"omitempty,min=4,max=20"`
	Name            string `json:"name" validate:"omitempty,min=4,max=20"`
	Surname         string `json:"surname" validate:"omitempty,min=4,max=20"`
	Email           string `json:"email" validate:"required,email"`
	Password        string `json:"password" validate:"required,min=4,max=20"`
	ConfirmPassword string `json:"confirmPassword" validate:"eqfield=Password"`
}

type UsersCreateResponse struct {
	User models.User `json:"user"`
}

type UsersCreate struct{}

func (uC *UsersCreate) Execute(data interface{}) (interface{}, error) {
	payload := data.(*UsersCreateRequest)
	userMapper := models.DefaultUserMapper
	emailSender := services.DefaultEmailSender

	existingUser, _ := userMapper.FindOne(sq.Eq{
		"email": payload.Email,
	})

	if existingUser.Email == payload.Email {
		return nil, &errors.Error{Status: 0, Reason: "Email already taken"}
	}

	user, err := userMapper.NewUser(
		payload.UserName,
		payload.Name,
		payload.Email,
		payload.Surname,
		payload.Password,
	)

	if err != nil {
		return nil, err
	}

	emailsToSend := []string{user.Email}
	templateData := map[string]string{
		"Name": user.UserName,
	}

	emailSender.Send(
		emailsToSend,
		"body",
		templateData,
	)

	dumpedUser := utils.DumpUser(user)

	return UsersCreateResponse{dumpedUser}, nil
}

func (uC *UsersCreate) Validate(data interface{}) error {
	err := validate.Struct(data)
	if err != nil {
		return &errors.Error{Status: 0, Reason: err.Error()}
	}
	return nil
}
