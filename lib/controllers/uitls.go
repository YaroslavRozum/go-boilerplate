package controllers

import (
	"encoding/json"
	"net/http"
	"reflect"

	"github.com/YaroslavRozum/go-boilerplate/lib/errors"
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

func defaultPayloadBuilder(payloadStruct interface{}) PayloadBuilder {
	return func(r *http.Request) (interface{}, error) {
		contentType := r.Header.Get("Content-Type")
		if contentType != "application/json" {
			return nil, &errors.Error{Status: 0, Reason: "NOT_JSON"}
		}
		decoder := json.NewDecoder(r.Body)
		plStrctT := reflect.TypeOf(payloadStruct)
		plStrctKind := plStrctT.Kind()
		var plStrctEl reflect.Type
		switch plStrctKind {
		case reflect.Struct:
			plStrctEl = plStrctT
		case reflect.Ptr:
			plStrctEl = plStrctT.Elem()
		}
		newStruct := reflect.New(plStrctEl)
		newStructIface := newStruct.Interface()
		err := decoder.Decode(newStructIface)
		if err != nil {
			return nil, &errors.Error{Status: 0, Reason: "WRONG_PAYLOAD"}
		}
		if plStrctKind == reflect.Struct {
			return newStruct.Elem().Interface(), nil
		}
		return newStructIface, nil
	}
}

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
