package factory

import (
	"net/http"

	"github.com/YaroslavRozum/go-boilerplate/lib"
	"github.com/YaroslavRozum/go-boilerplate/lib/controllers"
	"github.com/YaroslavRozum/go-boilerplate/lib/models"
	"github.com/YaroslavRozum/go-boilerplate/lib/router"
	"github.com/YaroslavRozum/go-boilerplate/lib/services"
	"github.com/YaroslavRozum/go-boilerplate/settings"
	"gopkg.in/go-playground/validator.v9"
)

func CreateServer(settings settings.Settings) (http.Server, error) {
	db, err := models.InitDB(settings)
	if err != nil {
		return http.Server{}, err
	}
	queryBuilder := models.QueryBuilder
	mappers := models.CreateMappers(queryBuilder, db)
	validate := validator.New()
	emailSender := lib.NewEmailSender(settings)
	services := services.CreateServices(settings, emailSender, validate, mappers)
	controller := controllers.CreateController(services)
	router := router.CreateRouter(controller)
	return http.Server{Addr: settings.Port, Handler: router}, nil
}
