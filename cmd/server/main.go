package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/andreyramos/go-metrics/internal/handlers"
	"github.com/andreyramos/go-metrics/internal/storage"
)

func main() {

	db := storage.NewMemStorage()

	r := chi.NewRouter()

	r.Route("/", func(r chi.Router) {
		r.Get("/", handlers.GetAll())
		r.Route("/value/{type}", func(r chi.Router) {
			r.Get("/{name}", handlers.Get())
		})
		r.Route("/update/{type}", func(r chi.Router) {
			r.Post("/{name}/{value}", handlers.Post(db))
		})
	})

	log.Fatal(http.ListenAndServe(":8080", r))

}
