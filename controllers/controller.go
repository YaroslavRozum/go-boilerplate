package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/YaroslavRozum/go-boilerplate/errors"
)

type Runner interface {
	Run(interface{}) (interface{}, error)
}

// RunnableContext -> context from r.Context, with which you can create a Runner (e.g you can store User in it)
type RunnableContext = interface{}
type CreateRunner func(RunnableContext) Runner

// PayloadBuilder must return a data from request or error if something wrong
type PayloadBuilder func(*http.Request) (interface{}, error)

// ResponseBuilder must write to ResponseWriter data that will be passed as second argument,
// data it is what Runner will return from Run method
type ResponseBuilder func(http.ResponseWriter, interface{})

func handleError(w http.ResponseWriter, err error) {
	w.Header().Set("Content-Type", "application/json")
	if _, ok := err.(*errors.Error); !ok {
		w.Write([]byte(`{"status":0, "reason":"Server Error" }`))
		return
	}
	jsonError, _ := json.Marshal(err)
	w.Write(jsonError)
	return
}

func NewController(cR CreateRunner, plB PayloadBuilder, rsB ResponseBuilder) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		payload, err := plB(r)
		if err != nil {
			handleError(w, err)
			return
		}
		ctx := r.Context().Value("context")
		serviceRunner := cR(ctx)
		result, err := serviceRunner.Run(payload)
		if err != nil {
			handleError(w, err)
			return
		}
		rsB(w, result)
	})
}
