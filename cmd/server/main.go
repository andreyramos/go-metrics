package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/andreyramos/go-metrics/internal/handlers"
	"github.com/andreyramos/go-metrics/internal/storage"
)

func main() {

	var flagAddr string
	flag.StringVar(&flagAddr, "a", "localhost:8080", "address and port to run agent")

	flag.Parse()

	db := storage.NewMemStorage()

	r := chi.NewRouter()

	r.Route("/", func(r chi.Router) {
		r.Get("/", handlers.GetAll(db))
		r.Route("/value/{type}", func(r chi.Router) {
			r.Get("/{name}", handlers.Get(db))
		})
		r.Route("/update/{type}", func(r chi.Router) {
			r.Post("/{name}/{value}", handlers.Post(db))
		})
	})

	log.Fatal(http.ListenAndServe(flagAddr, r))

}
