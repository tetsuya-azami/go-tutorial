package main

import (
	"mvc-api/controller"
	"mvc-api/domain"
	api "mvc-api/openapi"
	"mvc-api/repository"
	"net/http"

	"github.com/gorilla/mux"
)

var ItemRepository *repository.ItemRepository

func init() {
	mc := &domain.MyClock{}
	ItemRepository = repository.NewItemRepository(mc)
}

func main() {
	sh := api.NewStrictHandler(controller.NewServer(), nil)
	r := mux.NewRouter()
	h := api.HandlerFromMux(sh, r)

	http.ListenAndServe(":8080", h)
}
