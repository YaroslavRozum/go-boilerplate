package main

import (
	c "github.com/YaroslavRozum/go-boilerplate/controllers"
	"github.com/go-chi/chi"
)

func createRouter() *chi.Mux {
	controller := c.CreateController()
	r := chi.NewRouter()
	sessionCheck := controller.Sessions.Check
	r.With(sessionCheck).Get("/users", controller.Users.List)
	r.With(sessionCheck).Get("/products", controller.Products.List)
	r.Post("/sessions", controller.Sessions.Create)
	return r
}
