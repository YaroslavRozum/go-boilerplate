package lib

import (
	c "github.com/YaroslavRozum/go-boilerplate/lib/controllers"
	"github.com/go-chi/chi"
)

func CreateRouter() *chi.Mux {
	controller := c.CreateController()
	r := chi.NewRouter()
	sessionCheck := controller.Sessions.Check
	r.Post("/users", controller.Users.Create)
	r.With(sessionCheck).Get("/users", controller.Users.List)
	r.With(sessionCheck).Get("/products", controller.Products.List)
	r.Post("/sessions", controller.Sessions.Create)
	return r
}
