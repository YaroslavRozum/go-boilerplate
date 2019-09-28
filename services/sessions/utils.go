package sessions

import (
	"github.com/dgrijalva/jwt-go"
)

type Context struct {
	Email string `json:"email"`
	ID    string `json:"id"`
	Role  string `json:"role,omitempty"`
}

type Claims struct {
	Context Context `json:"context"`
	jwt.StandardClaims
}
