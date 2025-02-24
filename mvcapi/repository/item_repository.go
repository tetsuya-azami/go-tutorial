package repository

import (
	"fmt"
	"mvc-api/domain"
	"mvc-api/repository/internal/data"
)

type ItemRepositoryInterface interface {
	GetItems() []*domain.ItemRead
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
		itemRead, err := domain.NewItemRead(item.Id, item.JanCode, item.ItemName, item.Price, item.CategoryId, item.SeriesId, item.Stock, item.Discontinued, item.ReleaseDate, item.DeletedAt)
		if err != nil {
			fmt.Println("itemRead construction error")
		}
		itemReads = append(itemReads, itemRead)
	}

	return itemReads
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
