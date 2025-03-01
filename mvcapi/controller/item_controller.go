package controller

import (
	"context"
	"mvc-api/domain"
	api "mvc-api/openapi"
	"mvc-api/usecase"

	"github.com/oapi-codegen/runtime/types"
)

type ItemGetter interface {
	GetItems() []*domain.ItemRead
	GetItemById(id string) (*domain.ItemRead, usecase.ItemGetterErrors)
}

type ItemController struct {
	itemGetter ItemGetter
}

func NewItemsController(ig ItemGetter) *ItemController {
	return &ItemController{itemGetter: ig}
}

func (ic *ItemController) GetItems(ctx context.Context, request api.GetItemsRequestObject) (api.GetItemsResponseObject, error) {
	items := ic.itemGetter.GetItems()
	resp := api.GetItems200JSONResponse{Items: []api.Item{}}
	for _, item := range items {
		resp.Items = append(resp.Items, api.Item{
			Id:           item.Id(),
			ItemName:     item.ItemName(),
			JanCode:      item.JanCode(),
			Price:        int64(item.Price()),
			CategoryId:   int64(item.CategoryId()),
			SeriesId:     int64(item.SeriesId()),
			Stock:        int64(item.Stock()),
			Discontinued: item.Discontinued(),
			ReleasedDate: types.Date{Time: item.ReleaseDate()},
		})
	}

	return resp, nil
}
