package main

import (
	"log"
	"net/http"

	"github.com/YaroslavRozum/go-boilerplate/models"
	"github.com/YaroslavRozum/go-boilerplate/settings"
)

func main() {
	settings.InitSettings()
	defaultSettings := settings.DefaultSettings
	err := models.InitConn()
	if err != nil {
		log.Fatal(err.Error())
	}
	r := createRouter()
	log.Printf("Running server on Port %s", defaultSettings.Port)
	http.ListenAndServe(defaultSettings.Port, r)
}
