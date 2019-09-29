package main

import (
	"log"
	"net/http"

	"github.com/YaroslavRozum/go-boilerplate/models"
	"github.com/YaroslavRozum/go-boilerplate/services"
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
	r := createRouter()
	log.Printf("Running server on Port %s", defaultSettings.Port)
	http.ListenAndServe(defaultSettings.Port, r)
}
