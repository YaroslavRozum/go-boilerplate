package sessions

import (
	"fmt"

	"github.com/YaroslavRozum/go-boilerplate/errors"
	"github.com/YaroslavRozum/go-boilerplate/settings"
	"github.com/dgrijalva/jwt-go"
)

type SessionsCheck struct{}

func (s *SessionsCheck) Execute(auth string) (*Context, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(auth, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			errorMsg := fmt.Sprintf("unexpected signing Method: %v", token.Header["alg"])
			msg := &errors.Error{Status: 0, Reason: errorMsg}
			return nil, msg
		}
		return settings.DefaultSettings.JwtSecret, nil
	})
	if err != nil {
		return nil, &errors.Error{Status: 0, Reason: "Unauthorized"}
	}
	if token == nil || !token.Valid {
		return nil, &errors.Error{Status: 0, Reason: "Unauthorized"}
	}
	return &claims.Context, nil
}
