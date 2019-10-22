package controllers

import (
	"context"
	"net/http"

	"github.com/YaroslavRozum/go-boilerplate/lib/runner"
	"github.com/YaroslavRozum/go-boilerplate/lib/services/sessions"
	"github.com/YaroslavRozum/go-boilerplate/lib/services/utils"
)

type SessionsControllers struct {
	Check  func(http.Handler) http.Handler
	Create http.HandlerFunc
}

func CreateSessionsControllers(s sessions.Services) SessionsControllers {
	sessionCheck := s.Check()

	return SessionsControllers{
		Check: func(next http.Handler) http.Handler {
			return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				auth := r.Header.Get("Authorization")
				ctxPayload, err := sessionCheck.Execute(auth)
				if err != nil {
					utils.HandleError(w, err)
					return
				}
				ctx := context.WithValue(r.Context(), "context", ctxPayload)
				next.ServeHTTP(w, r.WithContext(ctx))
			})
		},
		Create: runner.NewController(
			runner.NewServiceRunnerCreator(s.Create),
			defaultPayloadBuilder(sessions.SessionsCreateRequest{}),
			defaultJsonResponse,
		),
	}
}
