package controllers

import "github.com/YaroslavRozum/go-boilerplate/lib/services"

type Controller struct {
	Users    UsersControllers
	Products ProductsControllers
	Sessions SessionsControllers
}

func CreateController(services services.Services) Controller {
	return Controller{
		Users:    CreateUsersControllers(services.Users),
		Products: CreateProductsControllers(services.Products),
		Sessions: CreateSessionsControllers(services.Sessions),
	}
}
