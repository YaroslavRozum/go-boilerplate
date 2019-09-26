package controllers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/YaroslavRozum/go-boilerplate/errors"
	"github.com/YaroslavRozum/go-boilerplate/services/sessions"
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
				ctxPayload, err := sessionCheck.Run(auth)
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
			func(r *http.Request) (interface{}, error) {
				contentType := r.Header.Get("Content-Type")
				if contentType != "application/json" {
					return nil, &errors.Error{Status: 0, Reason: "NOT_JSON"}
				}
				decoder := json.NewDecoder(r.Body)
				requestData := &sessions.SessionsCreateRequest{}
				err := decoder.Decode(requestData)
				if err != nil {
					return nil, &errors.Error{Status: 0, Reason: "WRONG_PAYLOAD"}
				}
				return requestData, nil
			},
			defaultJsonResponse,
		),
	}
}
