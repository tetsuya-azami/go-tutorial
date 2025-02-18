package main

import (
	"encoding/json"
	"fmt"
	"gorilla-tutorial/repository"
	"net/http"

	"strconv"

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

	r.HandleFunc("/users/{id}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			http.Error(w, "Invalid user ID", http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		user := repository.Users[id]
		if user == nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(`{"error": "User not found"}`))
			return
		}

		json, err := json.Marshal(user)
		if err != nil {
			fmt.Println(err)
			return
		}
		w.Write(json)
	})

	http.ListenAndServe(":8080", r)
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r)
		next.ServeHTTP(w, r)
	})
}
