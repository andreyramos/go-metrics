package handlers

import (
	"fmt"
	"net/http"

	"github.com/andreyramos/go-metrics/internal/storage"
	"github.com/go-chi/chi/v5"
)

func GetAll(db storage.Repositories) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		res.Write(db.ReadAll())
	}
}

func Get() http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		res.Write([]byte("GET"))
	}
}

func Post(db storage.Repositories) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		metricType := chi.URLParam(req, "type")
		name := chi.URLParam(req, "name")
		value := chi.URLParam(req, "value")

		body := "POST\r\n"
		body += metricType + "\r\n"
		body += name + "\r\n"
		body += value + "\r\n"

		switch metricType {
		case "gauge":
			var g storage.Guage
			err := g.FromString(value)
			if err != nil {
				msg := fmt.Sprintf("value %v not acceptable - %v", name, err)
				http.Error(res, msg, http.StatusBadRequest)
				return
			}
			db.SaveGuage(name, g)
		case "counter":
			var c storage.Counter
			err := c.FromString(value)
			if err != nil {
				msg := fmt.Sprintf("value %v not acceptable - %v", name, err)
				http.Error(res, msg, http.StatusBadRequest)
				return
			}
			db.SaveCounter(name, c)
		default:
			err := fmt.Errorf("not implemeted")
			http.Error(res, err.Error(), http.StatusNotImplemented)
			return
		}

		res.WriteHeader(http.StatusOK)
		res.Write([]byte(body))
	}
}
