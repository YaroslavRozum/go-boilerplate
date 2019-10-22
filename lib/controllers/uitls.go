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

func defaultPayloadBuilder(payloadStruct interface{}) runner.PayloadBuilder {
	return func(r *http.Request) (interface{}, error) {
		contentType := r.Header.Get("Content-Type")
		if contentType != "application/json" {
			return nil, &errors.Error{Status: 0, Reason: "NOT_JSON"}
		}
		decoder := json.NewDecoder(r.Body)
		payloadStrctT := reflect.TypeOf(payloadStruct)
		payloadStrctKind := payloadStrctT.Kind()
		var payloadStrctEl reflect.Type
		switch payloadStrctKind {
		case reflect.Struct:
			payloadStrctEl = payloadStrctT
		case reflect.Ptr:
			payloadStrctEl = payloadStrctT.Elem()
		}
		newStruct := reflect.New(payloadStrctEl)
		newStructIface := newStruct.Interface()
		err := decoder.Decode(newStructIface)
		if err != nil {
			return nil, &errors.Error{Status: 0, Reason: "WRONG_PAYLOAD"}
		}
		if payloadStrctKind == reflect.Struct {
			return newStruct.Elem().Interface(), nil
		}
		return newStructIface, nil
	}
}
