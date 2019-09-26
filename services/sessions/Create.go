package sessions

import (
	"github.com/YaroslavRozum/go-boilerplate/errors"
	"github.com/YaroslavRozum/go-boilerplate/services"
	"github.com/YaroslavRozum/go-boilerplate/settings"
	"github.com/dgrijalva/jwt-go"
)

var validate = services.Validate

type SessionsCreateRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type SessionsCreateResponse struct {
	Jwt string `json:"jwt"`
}

type SessionsCreate struct{}

func (s *SessionsCreate) Execute(data interface{}) interface{} {
	payload := data.(*SessionsCreateRequest)
	claims := &Claims{
		Context: Context{
			Email: payload.Email,
			ID:    "uuid",
			Role:  "ADMIN",
		},
		StandardClaims: jwt.StandardClaims{},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, _ := token.SignedString(settings.DefaultSettings.JwtSecret)

	return SessionsCreateResponse{tokenString}
}

func (pL *SessionsCreate) Validate(data interface{}) error {
	err := validate.Struct(data)
	if err != nil {
		return &errors.Error{Status: 0, Reason: err.Error()}
	}
	return nil
}
