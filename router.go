package main

import (
	"context"
	"net/http"

	"github.com/go-chi/chi"
)

func createRouter() *chi.Mux {
	var controllers = CreateControllers()
	r := chi.NewRouter()
	r.With(SessionCheck).Get("/users", controllers.UsersList)
	r.With(SessionCheck).Get("/products", controllers.ProductsList)
	return r
}

func SessionCheck(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		auth := r.Header.Get("Authorization")
		ctx := context.WithValue(r.Context(), "context", auth)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
