package sessions

import (
	"github.com/YaroslavRozum/go-boilerplate/lib/models"
	"github.com/YaroslavRozum/go-boilerplate/lib/runner"
	"gopkg.in/go-playground/validator.v9"
)

type Services struct {
	Create func() runner.Service
	Check  func() *SessionsCheck
}

func CreateSessionsServices(jwtSecret []byte, mappers models.Mappers, validate *validator.Validate) Services {
	return Services{
		Create: func() runner.Service {
			return &SessionsCreate{jwtSecret, mappers, validate}
		},
		Check: func() *SessionsCheck {
			return &SessionsCheck{mappers: mappers, jwtSecret: jwtSecret}
		},
	}
}
