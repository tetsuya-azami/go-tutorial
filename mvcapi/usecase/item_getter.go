package usecase

import (
	"fmt"
	"mvc-api/domain"
	"mvc-api/repository"
	"mvc-api/repository/rcustomerr"
	"mvc-api/usecase/ucustomerr"
)

type ItemGetterInterface interface {
	GetItems() []*domain.ItemRead
	GetItemById(id string) (*domain.ItemRead, ucustomerr.UsecaseErrorInterface)
}

type ItemGetter struct {
	itemRepository repository.ItemRepositoryInterface
}

func NewItemGetter(iri repository.ItemRepositoryInterface) *ItemGetter {
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

func (ig *ItemGetter) GetItemById(id string) (*domain.ItemRead, ucustomerr.UsecaseErrorInterface) {
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
