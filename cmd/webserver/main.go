package main

import (
	simplemetrics "github.com/co0p/simpleMetricsServiceInGo"
	"log"
	"net/http"
	"os"
)

const DefaultPort = "8080"

// Example implementation of a webservice collecting events
func main() {
	port := os.Getenv("port")
	if len(port) < 1 {
		port = DefaultPort
	}

	metricsService := simplemetrics.InMemorySimpleMetricsService{}

	collectHandler := func(res http.ResponseWriter, req *http.Request) {
		if req.Method != http.MethodGet {
			res.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		label := req.FormValue("label")
		if len(label) > 1 {
			res.WriteHeader(http.StatusBadRequest)
			return
		}

		e := simplemetrics.Event{Label: label}
		metricsService.Record(e)
	}

	aggregateHandler := func(res http.ResponseWriter, req *http.Request) {
		// TODO
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/collect", collectHandler)
	mux.HandleFunc("/aggregate", aggregateHandler)

	log.Printf("starting server on port: %s", port)
	if err := http.ListenAndServe(":"+port, mux); err != nil {
		log.Printf("failed to start server on port: %s", port)
	}
}
