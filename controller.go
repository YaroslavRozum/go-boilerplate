package main

import (
	"encoding/json"
	"net/http"

	"github.com/YaroslavRozum/go-boilerplate/errors"
)

type Runnable interface {
	Run(interface{}) (interface{}, error)
}

type CreateRunnable func(interface{}) Runnable

type PayloadBuilder func(*http.Request) interface{}

type ResponseBuilder func(http.ResponseWriter, interface{})

func NewController(cR CreateRunnable, plB PayloadBuilder, rsB ResponseBuilder) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		payload := plB(r)
		ctx := r.Context().Value("context")
		serviceRunner := cR(ctx)
		result, err := serviceRunner.Run(payload)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			if _, ok := err.(*errors.Error); !ok {
				w.Write([]byte(`{"status":0, "reason":"Server Error" }`))
				return
			}
			jsonError, _ := json.Marshal(err)
			w.Write(jsonError)
			return
		}
		rsB(w, result)
	})
}
