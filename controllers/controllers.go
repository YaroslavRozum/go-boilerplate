package controllers

import (
	"encoding/json"
	"net/http"
)

type Controller struct {
	Users    UsersControllers
	Products ProductsControllers
	Sessions SessionsControllers
}

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

func CreateController() Controller {
	return Controller{
		Users:    CreateUsersControllers(),
		Products: CreateProductsControllers(),
		Sessions: CreateSessionsControllers(),
	}
}
