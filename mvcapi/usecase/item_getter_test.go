package usecase

import (
	"mvc-api/domain"
	"mvc-api/repository"
	"mvc-api/repository/rcustomerr"
	"mvc-api/usecase/ucustomerr"
	"testing"
	"time"
)

var clock domain.Clock = &domain.FixedClock{
	FixedTime: time.Now(),
}

type expectedBehabior int

type mockedItemRepository struct {
	items            []*domain.ItemRead
	expectedBehabior expectedBehabior
}

const (
	Normal expectedBehabior = iota
	DataNotFound
	TooManyResults
)

func (mir *mockedItemRepository) GetItems() []*domain.ItemRead {
	return mir.items
}

func (mir *mockedItemRepository) GetItemById(id string) (*domain.ItemRead, repository.ItemRepositoryErrors) {
	switch mir.expectedBehabior {
	case Normal:
		return domain.NewItemRead("1", "327390283080", "item_1", 2500, 1, 1, 100, false, clock.Now(), time.Time{}), nil
	case DataNotFound:
		return nil, &rcustomerr.DataNotFoundError{}
	case TooManyResults:
		return nil, &rcustomerr.TooManyResultsFoundError{}
	}
	return nil, nil
}

func TestItemGetter_GetItems_結果が0件(t *testing.T) {
	// arrange
	itemGetter := NewItemGetter(&mockedItemRepository{items: []*domain.ItemRead{}})
	wantCount := 0

	// act
	actuals := itemGetter.GetItems()

	// assert
	if len(actuals) != wantCount {
		t.Errorf("len(actual) = %v, want %v", len(actuals), wantCount)
	}
}

func TestItemGetter_GetItems_結果が1件(t *testing.T) {
	// arrange
	itemRead1 := domain.NewItemRead("1", "327390283080", "item_1", 2500, 1, 1, 100, false, clock.Now(), time.Time{})
	itemGetter := NewItemGetter(&mockedItemRepository{items: []*domain.ItemRead{itemRead1}})
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

func TestItemGetter_GetItems_結果が複数件(t *testing.T) {
	// arrange
	itemRead1 := domain.NewItemRead("1", "327390283080", "item_1", 2500, 1, 1, 100, false, clock.Now(), time.Time{})
	itemRead2 := domain.NewItemRead("2", "222222222222", "item_2", 1200, 2, 2, 200, false, clock.Now(), time.Time{})
	items := []*domain.ItemRead{itemRead1, itemRead2}
	itemGetter := NewItemGetter(&mockedItemRepository{items: items})
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
	itemGetter := NewItemGetter(&mockedItemRepository{items: []*domain.ItemRead{itemRead}})
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
	itemGetter := NewItemGetter(&mockedItemRepository{expectedBehabior: DataNotFound})

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
	itemGetter := NewItemGetter(&mockedItemRepository{expectedBehabior: TooManyResults})

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
