package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

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

	http.ListenAndServe(":8080", r)
}
