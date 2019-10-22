package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/YaroslavRozum/go-boilerplate/lib/factory"
	"github.com/YaroslavRozum/go-boilerplate/settings"
)

func main() {
	defaultSettings := settings.CreateSettings()
	server, err := factory.CreateServer(defaultSettings)
	if err != nil {
		log.Fatal(err.Error())
	}

	idleConnsClosed := make(chan struct{})
	go func() {
		sigint := make(chan os.Signal)

		signal.Notify(sigint, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
		<-sigint

		log.Printf("Doing graceful shutdown")

		if err := server.Shutdown(context.Background()); err != nil {
			log.Printf("HTTP server Shutdown: %v", err)
		}
		close(idleConnsClosed)
	}()

	log.Printf("Running server on port %s", server.Addr)
	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		log.Printf("HTTP server ListenAndServe: %v", err)
	}

	<-idleConnsClosed
}
