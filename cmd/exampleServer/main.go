package main

import (
	simplemetrics "github.com/co0p/simpleMetricsServiceInGo/pkg"
	"log"
	"net/http"
	"os"
	"strconv"
)

const DefaultPort = "8080"

func main() {
	port := os.Getenv("port")
	if len(port) < 1 {
		port = DefaultPort
	}

	metricsService := simplemetrics.InMemorySimpleMetricsService{}

	collectHandler := func(res http.ResponseWriter, req *http.Request) {
		if req.Method != http.MethodGet {
			log.Printf("only GET requests allowed\n")
			http.Error(res, "only GET requests allowed", http.StatusMethodNotAllowed)
			return
		}

		// label is mandatory
		label := req.URL.Query().Get("label")
		if len(label) < 1 {
			log.Printf("missing 'label' in  query\n")
			http.Error(res, "missing 'label' in  query", http.StatusBadRequest)
			return
		}

		// value is optional, so don't abort
		valueStr := req.URL.Query().Get("value")
		value := 0
		if len(valueStr) > 0 {
			parsedVal, err := strconv.Atoi(valueStr)
			if err != nil {
				log.Printf("invalid value for  'value' in  query\n")
				http.Error(res, "invalid value for  'value' in  query", http.StatusBadRequest)
				return
			}
			value = parsedVal
		}

		e := simplemetrics.NewEvent(label, value)
		metricsService.Record(e)
		log.Printf("collected %s\n", e.String())
	}

	aggregateHandler := func(res http.ResponseWriter, req *http.Request) {
		// TODO
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/collect", collectHandler)
	mux.HandleFunc("/aggregate", aggregateHandler)

	log.Printf("** starting example metrics server on port: %s **\n", port)
	log.Printf(" - make a GET call to '/collect?label=<your event name>' in order to collect an event\n")
	log.Printf(" - make a GET call to '/aggregate/sum' in order to get the aggregation of your events\n")
	if err := http.ListenAndServe(":"+port, mux); err != nil {
		log.Printf("failed to start server on port: %s", port)
	}
}
