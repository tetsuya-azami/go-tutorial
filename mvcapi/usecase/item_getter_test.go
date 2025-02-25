package usecase

import (
	"mvc-api/domain"
	"testing"
	"time"
)

var clock domain.Clock = &domain.FixedClock{
	FixedTime: time.Now(),
}

type MockedItemRepository struct {
	items []*domain.ItemRead
}

func (mir *MockedItemRepository) GetItems() []*domain.ItemRead {
	return mir.items
}

func TestItemGetter_GetItem_NoItem(t *testing.T) {
	// arrange
	itemGetter := NewItemGetter(&MockedItemRepository{items: []*domain.ItemRead{}})
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
	itemRead1, _ := domain.NewItemRead("1", "327390283080", "item_1", 2500, 1, 1, 100, false, clock.Now(), time.Time{})
	itemGetter := NewItemGetter(&MockedItemRepository{items: []*domain.ItemRead{itemRead1}})
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
	itemRead1, _ := domain.NewItemRead("1", "327390283080", "item_1", 2500, 1, 1, 100, false, clock.Now(), time.Time{})
	itemRead2, _ := domain.NewItemRead("2", "222222222222", "item_2", 1200, 2, 2, 200, false, clock.Now(), time.Time{})
	items := []*domain.ItemRead{itemRead1, itemRead2}
	itemGetter := NewItemGetter(&MockedItemRepository{items: items})
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
