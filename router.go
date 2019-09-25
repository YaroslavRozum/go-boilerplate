package main

import (
	"context"
	"net/http"

	c "github.com/YaroslavRozum/go-boilerplate/controllers"
	"github.com/go-chi/chi"
)

func createRouter() *chi.Mux {
	controller := c.CreateController()
	r := chi.NewRouter()
	r.With(SessionCheck).Get("/users", controller.Users.List)
	r.With(SessionCheck).Get("/products", controller.Products.List)
	return r
}

func SessionCheck(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		auth := r.Header.Get("Authorization")
		ctx := context.WithValue(r.Context(), "context", auth)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
