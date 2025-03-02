package server

import (
	"mvc-api/controller"
	"mvc-api/domain"
	"mvc-api/repository"
	"mvc-api/usecase"
)

var ItemRepository *repository.ItemRepository
var ItemGetter *usecase.ItemGetter
var ItemsController *controller.ItemController

func init() {
	mc := &domain.MyClock{}
	ItemRepository = repository.NewItemRepository(mc)
	ItemGetter = usecase.NewItemGetter(ItemRepository)
	ItemsController = controller.NewItemsController(ItemGetter)
}

type Server struct {
	*controller.PingController
	*controller.ItemController
}

func NewServer() Server {
	return Server{
		PingController: &controller.PingController{},
		ItemController: ItemsController,
	}
}
