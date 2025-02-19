package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status": "ok"}`))
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/health", HealthCheckHandler)
	log.Fatal(http.ListenAndServe(":8080", r))
}
