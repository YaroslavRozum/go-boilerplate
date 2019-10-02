package controllers

type Controller struct {
	Users    UsersControllers
	Products ProductsControllers
	Sessions SessionsControllers
}

func CreateController() Controller {
	return Controller{
		Users:    CreateUsersControllers(),
		Products: CreateProductsControllers(),
		Sessions: CreateSessionsControllers(),
	}
}
