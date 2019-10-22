package sessions

import (
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/YaroslavRozum/go-boilerplate/lib/errors"
	"github.com/YaroslavRozum/go-boilerplate/lib/models"
	"github.com/dgrijalva/jwt-go"
)

type SessionsCheck struct {
	jwtSecret []byte
	mappers   models.Mappers
}

func (s *SessionsCheck) Execute(auth string) (Context, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(auth, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			errorMsg := fmt.Sprintf("unexpected signing Method: %v", token.Header["alg"])
			msg := &errors.Error{Status: 0, Reason: errorMsg}
			return nil, msg
		}
		return s.jwtSecret, nil
	})
	if err != nil {
		return Context{}, &errors.Error{Status: 0, Reason: "Unauthorized"}
	}
	ctx := claims.Context
	userMapper := s.mappers.UserMapper
	user, err := userMapper.FindOne(sq.Eq{
		"id":    ctx.ID,
		"email": ctx.Email,
	})
	if err != nil {
		return Context{}, &errors.Error{Status: 0, Reason: "Unauthorized"}
	}
	if user.ID != ctx.ID && user.Email != ctx.Email {
		return Context{}, &errors.Error{Status: 0, Reason: "Unauthorized"}
	}
	if token == nil || !token.Valid {
		return Context{}, &errors.Error{Status: 0, Reason: "Unauthorized"}
	}
	return claims.Context, nil
}
