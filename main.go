package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/ImanAski/todo-list-go/handlers"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

const port string = ":9000"

func main() {
	stopChan := make(chan os.Signal)
	signal.Notify(stopChan, os.Interrupt)
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", handlers.HomeHandler)
	r.Mount("/todo", handlers.Todohandler())

	srv := &http.Server{
		Addr:         port,
		Handler:      r,
		ReadTimeout:  60 * time.Second,
		WriteTimeout: 60 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	go func() {
		log.Println("Starting Server on port", port, "...")
		if err := srv.ListenAndServe(); err != nil {
			log.Println("There was a problem starting the server")
		}
	}()

	<-stopChan
	log.Println("Shutting down the server...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	srv.Shutdown(ctx)
	defer cancel()

}
