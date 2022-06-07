package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"goliveserver/internal/handler"
)

func main() {

	h := handler.NewHandler()

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", h.Index)
	log.Println("Starting web server...")

	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Stop web server...")
}
