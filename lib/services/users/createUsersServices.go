package users

import (
	"github.com/YaroslavRozum/go-boilerplate/lib"
	"github.com/YaroslavRozum/go-boilerplate/lib/models"
	"github.com/YaroslavRozum/go-boilerplate/lib/runner"
	"gopkg.in/go-playground/validator.v9"
)

type Services struct {
	Create func() runner.Service
	List   func() runner.Service
}

func CreateUserServices(mappers models.Mappers, emailSender lib.EmailSender, validate *validator.Validate) Services {
	return Services{
		Create: func() runner.Service {
			return &UsersCreate{mappers, emailSender, validate}
		},
		List: func() runner.Service {
			return &UsersList{mappers: mappers, validate: validate}
		},
	}
}
