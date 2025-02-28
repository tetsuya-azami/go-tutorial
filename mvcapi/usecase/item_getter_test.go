package usecase

import (
	"mvc-api/domain"
	"mvc-api/repository/rcustomerr"
	"mvc-api/usecase/ucustomerr"
	"testing"
	"time"
)

var clock domain.Clock = &domain.FixedClock{
	FixedTime: time.Now(),
}

type NormalMockedItemRepository struct {
	items []*domain.ItemRead
}

func (mir *NormalMockedItemRepository) GetItems() []*domain.ItemRead {
	return mir.items
}

func (mir *NormalMockedItemRepository) GetItemById(id string) (*domain.ItemRead, rcustomerr.RepositoryErrorInterface) {
	return domain.NewItemRead("1", "327390283080", "item_1", 2500, 1, 1, 100, false, clock.Now(), time.Time{}), nil
}

type DataNotFoundItemRepository struct {
}

func (mir *DataNotFoundItemRepository) GetItems() []*domain.ItemRead {
	return []*domain.ItemRead{}
}

func (mir *DataNotFoundItemRepository) GetItemById(id string) (*domain.ItemRead, rcustomerr.RepositoryErrorInterface) {
	return nil, &rcustomerr.DataNotFoundError{}
}

type TooManyResultsItemRepository struct {
}

func (mir *TooManyResultsItemRepository) GetItems() []*domain.ItemRead {
	return nil
}

func (mir *TooManyResultsItemRepository) GetItemById(id string) (*domain.ItemRead, rcustomerr.RepositoryErrorInterface) {
	return nil, &rcustomerr.TooManyResultsFoundError{}
}

func TestItemGetter_GetItem_NoItem(t *testing.T) {
	// arrange
	itemGetter := NewItemGetter(&NormalMockedItemRepository{items: []*domain.ItemRead{}})
	wantCount := 0

	// act
	actuals := itemGetter.GetItems()

	// assert
	if len(actuals) != wantCount {
		t.Errorf("len(actual) = %v, want %v", len(actuals), 1)
	}
}

func TestItemGetter_GetItem_OneItem(t *testing.T) {
	// arrange
	itemRead1 := domain.NewItemRead("1", "327390283080", "item_1", 2500, 1, 1, 100, false, clock.Now(), time.Time{})
	itemGetter := NewItemGetter(&NormalMockedItemRepository{items: []*domain.ItemRead{itemRead1}})
	wantCount := 1

	// act
	actuals := itemGetter.GetItems()

	// assert
	if len(actuals) != wantCount {
		t.Errorf("len(actual) = %v, want %v", len(actuals), 1)
	}
	if *actuals[0] != *itemRead1 {
		t.Errorf("actuals[0] = %v, want %v", *actuals[0], *itemRead1)
	}
}

func TestItemGetter_GetItem_MultipleItem(t *testing.T) {
	// arrange
	itemRead1 := domain.NewItemRead("1", "327390283080", "item_1", 2500, 1, 1, 100, false, clock.Now(), time.Time{})
	itemRead2 := domain.NewItemRead("2", "222222222222", "item_2", 1200, 2, 2, 200, false, clock.Now(), time.Time{})
	items := []*domain.ItemRead{itemRead1, itemRead2}
	itemGetter := NewItemGetter(&NormalMockedItemRepository{items: items})
	wantCount := 2

	// act
	actuals := itemGetter.GetItems()

	// assert
	if len(actuals) != wantCount {
		t.Errorf("len(actual) = %v, want %v", len(actuals), 1)
	}

	for i, item := range items {
		if *actuals[i] != *item {
			t.Errorf("actuals[%d] = %v, want %v", i, actuals[i], item)
		}
	}
}

func TestItemGetter_GetItemById_正常系(t *testing.T) {
	// arrange
	itemRead := domain.NewItemRead("1", "327390283080", "item_1", 2500, 1, 1, 100, false, clock.Now(), time.Time{})
	itemGetter := NewItemGetter(&NormalMockedItemRepository{items: []*domain.ItemRead{itemRead}})
	wantItem := itemRead

	// act
	actualItem, err := itemGetter.GetItemById("1")

	// assert
	if err != nil {
		t.Errorf("err = %v, do not want err", err)
	}
	if *actualItem != *wantItem {
		t.Errorf("item = %v, want %v", *actualItem, *wantItem)
	}
}

func TestItemGetter_GetItemById_対象のIdがない(t *testing.T) {
	// arrange
	itemGetter := NewItemGetter(&DataNotFoundItemRepository{})

	// act
	actualItem, err := itemGetter.GetItemById("2")

	// assert
	if actualItem != nil {
		t.Errorf("item = %v, want nil", actualItem)
	}
	if res, ok := err.(*ucustomerr.DataNotFoundError); !ok {
		t.Errorf("err = %v, want DataNotFoundError", res)
	}
}

func TestItemGetter_GetItemById_複数件ヒット(t *testing.T) {
	// arrange
	itemGetter := NewItemGetter(&TooManyResultsItemRepository{})

	// act
	actualItem, err := itemGetter.GetItemById("1")

	// assert
	if actualItem != nil {
		t.Errorf("item = %v, want nil", actualItem)
	}
	if res, ok := err.(*ucustomerr.TooManyResultsFoundError); !ok {
		t.Errorf("err = %v, want TooManyResultsFoundError", res)
	}
}
