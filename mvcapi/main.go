package main

import (
	"mvc-api/controller"
	"mvc-api/domain"
	api "mvc-api/openapi"
	"mvc-api/repository"
	"mvc-api/usecase"
	"net/http"

	"github.com/gorilla/mux"
)

var ItemRepository *repository.ItemRepository
var ItemGetter *usecase.ItemGetter

func init() {
	mc := &domain.MyClock{}
	ItemRepository = repository.NewItemRepository(mc)
	ItemGetter = usecase.NewItemGetter(ItemRepository)
}

func main() {
	sh := api.NewStrictHandler(controller.NewServer(), nil)
	r := mux.NewRouter()
	h := api.HandlerFromMux(sh, r)

	http.ListenAndServe(":8080", h)
}
