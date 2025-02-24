package usecase

import (
	"fmt"
	"mvc-api/domain"
	"mvc-api/repository"
)

type ItemGetter struct {
	itemRepository repository.ItemRepositoryInterface
}

func NewItemGetter(iri repository.ItemRepositoryInterface) *ItemGetter {
	return &ItemGetter{itemRepository: iri}
}

func (ig *ItemGetter) GetItem() []*domain.ItemRead {
	items := ig.itemRepository.GetItems()
	if len(items) == 0 {
		fmt.Println("no items")
		return make([]*domain.ItemRead, 0)
	}

	return items
}
