package controllers

import (
	"context"
	"net/http"

	"github.com/YaroslavRozum/go-boilerplate/lib/services/sessions"
)

type SessionsControllers struct {
	Check  func(http.Handler) http.Handler
	Create http.HandlerFunc
}

func CreateSessionsControllers() SessionsControllers {
	sessionCheck := &sessions.SessionsCheck{}

	return SessionsControllers{
		Check: func(next http.Handler) http.Handler {
			return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				auth := r.Header.Get("Authorization")
				ctxPayload, err := sessionCheck.Execute(auth)
				if err != nil {
					handleError(w, err)
					return
				}
				ctx := context.WithValue(r.Context(), "context", ctxPayload)
				next.ServeHTTP(w, r.WithContext(ctx))
			})
		},
		Create: NewController(
			NewServiceRunnerCreator(&sessions.SessionsCreate{}),
			defaultPayloadBuilder(&sessions.SessionsCreateRequest{}),
			defaultJsonResponse,
		),
	}
}
