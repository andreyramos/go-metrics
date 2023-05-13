package handlers

import (
	"fmt"
	"net/http"

	"github.com/andreyramos/go-metrics/internal/storage"
	"github.com/go-chi/chi/v5"
)

func GetAll(db storage.Repositories) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		gs, cs, err := db.GetAll()

		if err != nil {
			http.Error(res, err.Error(), http.StatusNotFound)
			return
		}

		body := "<h1>All metrics</h1>"
		body += "<h2>Guages</h2>"
		for k, v := range gs {
			body += fmt.Sprintf("<div>%s: %v</div>", k, v)
		}
		body += "<h2>Counters</h2>"
		for k, v := range cs {
			body += fmt.Sprintf("<div>%s: %v</div>", k, v)
		}

		res.Header().Set("content-type", "text/HTML")
		res.WriteHeader(http.StatusOK)
		res.Write([]byte(body))
	}
}

func Get(db storage.Repositories) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		metricType := chi.URLParam(req, "type")
		name := chi.URLParam(req, "name")
		var val string

		switch metricType {
		case "gauge":
			g, err := db.GetGauge(name)
			if err != nil {
				http.Error(res, err.Error(), http.StatusNotFound)
				return
			}
			val = fmt.Sprintf("%f", g)
		case "counter":
			c, err := db.GetCounter(name)
			if err != nil {
				http.Error(res, err.Error(), http.StatusNotFound)
				return
			}
			val = fmt.Sprintf("%d", c)
		default:
			err := fmt.Errorf("not implemeted")
			http.Error(res, err.Error(), http.StatusNotImplemented)
			return
		}

		res.WriteHeader(http.StatusOK)
		res.Write([]byte(val))
	}
}

func Post(db storage.Repositories) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		metricType := chi.URLParam(req, "type")
		name := chi.URLParam(req, "name")
		value := chi.URLParam(req, "value")

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
	}
}
