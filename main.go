package main

import (
	"log"
	"net/http"

	"github.com/YaroslavRozum/go-boilerplate/lib"
	"github.com/YaroslavRozum/go-boilerplate/lib/models"
	"github.com/YaroslavRozum/go-boilerplate/lib/services"
	"github.com/YaroslavRozum/go-boilerplate/settings"
)

func main() {
	settings.InitSettings()
	defaultSettings := settings.DefaultSettings
	err := models.InitDB()
	if err != nil {
		log.Fatal(err.Error())
	}
	models.InitModels()
	services.InitEmailSender()
	r := lib.CreateRouter()
	log.Printf("Running server on Port %s", defaultSettings.Port)
	http.ListenAndServe(defaultSettings.Port, r)
}
