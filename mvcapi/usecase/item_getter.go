package usecase

import (
	"fmt"
	"mvc-api/domain"
	"mvc-api/repository"
	"mvc-api/repository/rcustomerr"
	"mvc-api/usecase/ucustomerr"
)

type ItemRepositoryInterface interface {
	GetItems() []*domain.ItemRead
	GetItemById(id string) (*domain.ItemRead, repository.ItemRepositoryErrors)
}

type ItemGetterErrors interface {
	Error() string
}

type ItemGetter struct {
	itemRepository ItemRepositoryInterface
}

func NewItemGetter(iri ItemRepositoryInterface) *ItemGetter {
	return &ItemGetter{itemRepository: iri}
}

func (ig *ItemGetter) GetItems() []*domain.ItemRead {
	items := ig.itemRepository.GetItems()
	if len(items) == 0 {
		fmt.Println("no items")
		return make([]*domain.ItemRead, 0)
	}

	return items
}

func (ig *ItemGetter) GetItemById(id string) (*domain.ItemRead, ItemGetterErrors) {
	item, err := ig.itemRepository.GetItemById(id)
	if err != nil {
		if res, ok := err.(*rcustomerr.DataNotFoundError); ok {
			return nil, &ucustomerr.DataNotFoundError{Msg: res.Error(), Err: res}
		} else if res, ok := err.(*rcustomerr.TooManyResultsFoundError); ok {
			return nil, &ucustomerr.TooManyResultsFoundError{Msg: res.Error(), Err: res}
		}
	}

	return item, nil
}
