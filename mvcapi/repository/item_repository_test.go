package repository

import (
	"fmt"
	"mvc-api/domain"
	"mvc-api/repository/rcustomerr"
	"testing"
	"time"
)

var now time.Time

func Test_GetItems_正常系(t *testing.T) {
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

	itemRead1 := domain.NewItemRead("1", "327390283080", "item_1", 2500, 1, 1, 100, false, now, time.Time{})
	itemRead2 := domain.NewItemRead("2", "3273902878656", "item_2", 1200, 2, 2, 200, false, now, now)
	expected := []*domain.ItemRead{itemRead1, itemRead2}

	for i, item := range items {
		if *item != *expected[i] {
			t.Errorf("GetItems() = %v, want %v", item, expected[i])
		}
	}
}

func TestGetItemById_正常系(t *testing.T) {
	// arrange
	now = time.Now()
	fc := &domain.FixedClock{
		FixedTime: now,
	}
	ir := NewItemRepository(fc)

	item, _ := ir.GetItemById("1")
	expectedItem := domain.NewItemRead("1", "327390283080", "item_1", 2500, 1, 1, 100, false, now, time.Time{})

	if *item != *expectedItem {
		t.Errorf("GetItemById() = %v, want %v", item, expectedItem)
	}
}

func TestGetItemById_存在しないIdでの検索(t *testing.T) {
	// arrange
	now = time.Now()
	fc := &domain.FixedClock{
		FixedTime: now,
	}
	ir := NewItemRepository(fc)

	notExistId := "notExistId"
	expectedErrMsg := fmt.Sprintf("item not found by id: %v", notExistId)

	// act
	item, err := ir.GetItemById(notExistId)

	// assert
	if err != nil {
		if result, ok := err.(*rcustomerr.DataNotFoundError); ok {
			if result.Error() != expectedErrMsg {
				t.Errorf("want: %v, but: %v", expectedErrMsg, result)
			}
		} else {
			t.Errorf("error = %v, want DataNotFoundError", result)
		}
	}
	if item != nil {
		t.Errorf("item = %v, want nil", item)
	}
}

// TODO: 複数件取得される場合のテスト
