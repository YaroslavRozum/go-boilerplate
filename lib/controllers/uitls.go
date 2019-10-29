package controllers

import (
	"encoding/json"
	"net/http"
	"reflect"

	"github.com/YaroslavRozum/go-boilerplate/lib/errors"
	"github.com/YaroslavRozum/go-boilerplate/lib/runner"
)

type response struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
}

func defaultJsonResponse(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	res := response{
		Status: 1,
		Data:   data,
	}
	encoder := json.NewEncoder(w)
	encoder.Encode(res)
}

type payloadCreator func() interface{}

func defaultPayloadBuilder(createPayload payloadCreator) runner.PayloadBuilder {
	return func(r *http.Request) (interface{}, error) {
		contentType := r.Header.Get("Content-Type")
		if contentType != "application/json" {
			return nil, &errors.Error{Status: 0, Reason: "NOT_JSON"}
		}
		decoder := json.NewDecoder(r.Body)
		payload := createPayload()
		var err error
		payloadValue := reflect.ValueOf(payload)
		if payloadValue.Type().Kind() == reflect.Struct {
			vp := reflect.New(payloadValue.Type())
			vp.Elem().Set(payloadValue)
			err = decoder.Decode(vp.Interface())
			payload = vp.Elem().Interface()
		} else {
			err = decoder.Decode(payload)
		}
		if err != nil {
			return nil, &errors.Error{Status: 0, Reason: "WRONG_PAYLOAD"}
		}
		return payload, nil
	}
}
