package controllers

import (
	"encoding/json"
	"net/http"
)

type Controller struct {
	Users    UsersControllers
	Products ProductsControllers
}

func defaultJsonResponse(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	encoder.Encode(data)
}

func CreateController() Controller {
	return Controller{
		Users:    CreateUsersControllers(),
		Products: CreateProductsControllers(),
	}
}
