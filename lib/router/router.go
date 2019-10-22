package router

import (
	"github.com/YaroslavRozum/go-boilerplate/lib/controllers"
	"github.com/go-chi/chi"
)

func CreateRouter(controller controllers.Controller) *chi.Mux {
	r := chi.NewRouter()
	sessionCheck := controller.Sessions.Check
	r.Post("/users", controller.Users.Create)
	r.With(sessionCheck).Get("/users", controller.Users.List)
	r.With(sessionCheck).Get("/products", controller.Products.List)
	r.Post("/sessions", controller.Sessions.Create)
	return r
}
