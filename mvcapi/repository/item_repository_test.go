package repository

import (
	"mvc-api/domain"
	"testing"
	"time"
)

var now time.Time

func TestItemRepository_GetItems(t *testing.T) {
	// arrange
	now = time.Now()
	fc := &domain.FixedClock{
		FixedTime: now,
	}
	ir := NewItemRepository(fc)
	wantCount := 2

	// act
	items := ir.GetItems()

	// assert
	if len(items) != wantCount {
		t.Errorf("GetItems() = %v, want %v", len(items), wantCount)
	}

	expected := getExpectedItems(t)

	for i, item := range items {
		if *item != *expected[i] {
			t.Errorf("GetItems() = %v, want %v", item, expected[i])
		}
	}
}

func getExpectedItems(t *testing.T) []*domain.ItemRead {
	itemRead1, err := domain.NewItemRead("1", "327390283080", "item_1", 2500, 1, 1, 100, false, now, time.Time{})
	if err != nil {
		t.Fatal(err)
	}
	itemRead2, err := domain.NewItemRead("2", "3273902878656", "item_2", 1200, 2, 2, 200, false, now, now)
	if err != nil {
		t.Fatal(err)
	}

	return []*domain.ItemRead{itemRead1, itemRead2}
}
