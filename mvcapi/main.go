package main

import (
	"mvc-api/controller"
	api "mvc-api/openapi"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	sh := api.NewStrictHandler(controller.NewServer(), nil)
	r := mux.NewRouter()
	h := api.HandlerFromMux(sh, r)

	http.ListenAndServe(":8080", h)
}
