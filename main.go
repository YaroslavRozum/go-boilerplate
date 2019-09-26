package main

import (
	"net/http"

	"github.com/YaroslavRozum/go-boilerplate/settings"
)

func main() {
	settings.InitSettings()
	r := createRouter()
	http.ListenAndServe(settings.DefaultSettings.Port, r)
}
