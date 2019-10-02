package controllers

import (
	"net/http"
)

type Runner interface {
	Run(interface{}) (interface{}, error)
}

// RunnerContext -> context from r.Context, with which you can create a Runner (e.g you can store User in it)
type RunnerContext = interface{}
type CreateRunner func(RunnerContext) Runner

// PayloadBuilder must return a data from request or error if something wrong
type PayloadBuilder func(*http.Request) (interface{}, error)

// ResponseBuilder must write to ResponseWriter data that will be passed as second argument,
// data it is what Runner will return from Run method
type ResponseBuilder func(http.ResponseWriter, interface{})

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
