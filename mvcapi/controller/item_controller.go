package controller

import (
	"context"
	"mvc-api/domain"
	api "mvc-api/openapi"
	"mvc-api/usecase"
	"mvc-api/usecase/ucustomerr"

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
	respItems := ConvertToResponses(items)

	return api.GetItems200JSONResponse{Items: respItems}, nil
}

func (ic *ItemController) GetItemsById(ctx context.Context, request api.GetItemsByIdRequestObject) (api.GetItemsByIdResponseObject, error) {
	itemId := request.ItemId
	if itemId == "" {
		return api.GetItemsById404JSONResponse{Message: "itemが見つかりませんでした。"}, nil
	}
	item, err := ic.itemGetter.GetItemById(itemId)
	if res, ok := err.(*ucustomerr.DataNotFoundError); ok {
		return api.GetItemsById404JSONResponse{Message: res.Msg}, nil
	} else if res, ok := err.(*ucustomerr.TooManyResultsFoundError); ok {
		return api.GetItemsById500JSONResponse{Message: res.Msg}, nil
	}

	return api.GetItemsById200JSONResponse{
		Item: ConvertToResponse(item),
	}, nil
}

func ConvertToResponses(items []*domain.ItemRead) []api.Item {
	results := []api.Item{}
	for _, item := range items {
		results = append(results, ConvertToResponse(item))
	}

	return results
}

func ConvertToResponse(item *domain.ItemRead) api.Item {
	return api.Item{
		Id:           item.Id(),
		ItemName:     item.ItemName(),
		JanCode:      item.JanCode(),
		Price:        int64(item.Price()),
		CategoryId:   int64(item.CategoryId()),
		SeriesId:     int64(item.SeriesId()),
		Stock:        int64(item.Stock()),
		Discontinued: item.Discontinued(),
		ReleasedDate: types.Date{Time: item.ReleaseDate()},
	}
}
