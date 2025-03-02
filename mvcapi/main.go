package main

import (
	api "mvc-api/openapi"
	"mvc-api/server"
	"net/http"
)

func main() {
	sh := api.NewStrictHandler(server.NewServer(), nil)
	h := api.Handler(sh)

	http.ListenAndServe(":8080", h)
}
