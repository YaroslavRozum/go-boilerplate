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
		plStrctEl := reflect.TypeOf(payloadStruct).Elem()
		requestData := reflect.New(plStrctEl).Interface()
		err := decoder.Decode(requestData)
		if err != nil {
			return nil, &errors.Error{Status: 0, Reason: "WRONG_PAYLOAD"}
		}
		return requestData, nil
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
