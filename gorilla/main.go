package main

import (
	"fmt"
	"gorilla-tutorial/presentation/handler"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.Use(loggingMiddleware)

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		fmt.Println(vars)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello, World!"))
	})

	r.HandleFunc("/products", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		fmt.Println(vars)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello, Products!"))
	})

	r.Host("www.example.com").Path("/products/{id:[0-9]{0,9}}").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		fmt.Println(vars)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello, Products " + vars["id"] + "!"))
	})

	r.HandleFunc("/redirects", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Location", "https://www.google.com")
		w.WriteHeader(http.StatusFound)
	})

	r.HandleFunc("/users/{id}", handler.UsersHandler)

	http.ListenAndServe(":8080", r)
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r)
		next.ServeHTTP(w, r)
	})
}
