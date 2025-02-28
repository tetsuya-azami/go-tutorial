package repository

import (
	"fmt"
	"mvc-api/domain"
	"mvc-api/repository/internal/data"
	"mvc-api/repository/rcustomerr"
)

type ItemRepositoryInterface interface {
	GetItems() []*domain.ItemRead
	GetItemById(id string) (*domain.ItemRead, rcustomerr.RepositoryErrorInterface)
}

type ItemRepository struct {
	clock domain.Clock
}

func NewItemRepository(clock domain.Clock) *ItemRepository {
	return &ItemRepository{
		clock: clock,
	}
}

func (ir *ItemRepository) GetItems() []*domain.ItemRead {
	items := getItemData(ir.clock)
	itemReads := []*domain.ItemRead{}
	for _, item := range items {
		itemRead := domain.NewItemRead(item.Id, item.JanCode, item.ItemName, item.Price, item.CategoryId, item.SeriesId, item.Stock, item.Discontinued, item.ReleaseDate, item.DeletedAt)
		itemReads = append(itemReads, itemRead)
	}

	return itemReads
}

func (ir *ItemRepository) GetItemById(id string) (*domain.ItemRead, rcustomerr.RepositoryErrorInterface) {
	items := getItemData(ir.clock)
	var filtered []*data.ItemData
	for _, item := range items {
		if item.Id == id {
			filtered = append(filtered, item)
		}
	}

	if len(filtered) == 0 {
		err := fmt.Errorf("item not found by id: %v", id)
		return nil, &rcustomerr.DataNotFoundError{Msg: err.Error(), Err: err}
	} else if len(filtered) != 1 {
		err := fmt.Errorf("duplicate item found by id: %v", id)
		return nil, &rcustomerr.TooManyResultsFoundError{Msg: err.Error(), Err: err}
	}
	item := filtered[0]

	itemRead := domain.NewItemRead(item.Id, item.JanCode, item.ItemName, item.Price, item.CategoryId, item.SeriesId, item.Stock, item.Discontinued, item.ReleaseDate, item.DeletedAt)

	return itemRead, nil
}

func getItemData(clock domain.Clock) []*data.ItemData {
	return []*data.ItemData{
		{
			Id:           "1",
			JanCode:      "327390283080",
			ItemName:     "item_1",
			Price:        2500,
			CategoryId:   1,
			SeriesId:     1,
			Stock:        100,
			Discontinued: false,
			ReleaseDate:  clock.Now(),
			CreatedAt:    clock.Now(),
			UpdatedAt:    clock.Now(),
		},
		{
			Id:           "2",
			JanCode:      "3273902878656",
			ItemName:     "item_2",
			Price:        1200,
			CategoryId:   2,
			SeriesId:     2,
			Stock:        200,
			Discontinued: false,
			ReleaseDate:  clock.Now(),
			CreatedAt:    clock.Now(),
			UpdatedAt:    clock.Now(),
			DeletedAt:    clock.Now(),
		},
	}
}
