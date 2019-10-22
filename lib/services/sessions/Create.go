package sessions

import (
	sq "github.com/Masterminds/squirrel"
	"github.com/YaroslavRozum/go-boilerplate/lib/errors"
	"github.com/YaroslavRozum/go-boilerplate/lib/models"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/go-playground/validator.v9"
)

type SessionsCreateRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type SessionsCreateResponse struct {
	Jwt string `json:"jwt"`
}

type SessionsCreate struct {
	jwtSecret []byte
	mappers   models.Mappers
	validate  *validator.Validate
}

func (s *SessionsCreate) Execute(data interface{}) (interface{}, error) {
	payload := data.(SessionsCreateRequest)
	userMapper := s.mappers.UserMapper
	user, _ := userMapper.FindOne(sq.Eq{
		"email": payload.Email,
	})

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(payload.Password))

	if err != nil {
		return nil, &errors.Error{Status: 0, Reason: "Wrong email or password"}
	}

	claims := Claims{
		Context: Context{
			Email: user.Email,
			ID:    user.ID,
		},
		StandardClaims: jwt.StandardClaims{},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(s.jwtSecret)

	if err != nil {
		return nil, &errors.Error{Status: 0, Reason: err.Error()}
	}

	return SessionsCreateResponse{tokenString}, nil
}

func (pL *SessionsCreate) Validate(data interface{}) error {
	err := pL.validate.Struct(data)
	if err != nil {
		return &errors.Error{Status: 0, Reason: err.Error()}
	}
	return nil
}
