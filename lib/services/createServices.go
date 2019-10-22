package services

import (
	"github.com/YaroslavRozum/go-boilerplate/lib"
	"github.com/YaroslavRozum/go-boilerplate/lib/models"
	"github.com/YaroslavRozum/go-boilerplate/lib/services/products"
	"github.com/YaroslavRozum/go-boilerplate/lib/services/sessions"
	"github.com/YaroslavRozum/go-boilerplate/lib/services/users"
	"github.com/YaroslavRozum/go-boilerplate/settings"
	"gopkg.in/go-playground/validator.v9"
)

type Services struct {
	Users    users.Services
	Sessions sessions.Services
	Products products.Services
}

func CreateServices(
	settings settings.Settings,
	emailSender lib.EmailSender,
	validate *validator.Validate,
	mappers models.Mappers,
) Services {
	return Services{
		Users:    users.CreateUserServices(mappers, emailSender, validate),
		Sessions: sessions.CreateSessionsServices(settings.JwtSecret, mappers, validate),
		Products: products.CreateProductsServices(mappers, validate),
	}
}
