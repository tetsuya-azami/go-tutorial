package controller_test

import (
	"context"
	"mvc-api/controller"
	"mvc-api/domain"
	api "mvc-api/openapi"
	"mvc-api/usecase"
	"mvc-api/usecase/ucustomerr"
	"testing"
	"time"

	"github.com/oapi-codegen/runtime/types"
)

var now = time.Now()

var testItems = []*domain.ItemRead{
	domain.NewItemRead("1", "1234567890123", "item1", 100, 1, 1, 10, false, now, time.Time{}),
	domain.NewItemRead("2", "1234567890124", "item2", 200, 2, 2, 20, false, now, time.Time{}),
}

var expectedItems = []api.Item{
	{
		CategoryId:   1,
		Discontinued: false,
		Id:           "1",
		ItemName:     "item1",
		JanCode:      "1234567890123",
		Price:        100,
		ReleasedDate: types.Date{Time: now},
		SeriesId:     1,
		Stock:        10,
	},
	{
		CategoryId:   2,
		Discontinued: false,
		Id:           "2",
		ItemName:     "item2",
		JanCode:      "1234567890124",
		Price:        200,
		ReleasedDate: types.Date{Time: now},
		SeriesId:     2,
		Stock:        20,
	},
}

type ExpectedBehavior int

const (
	Normal ExpectedBehavior = iota
	NotFound
	TooManyResultsFound
)

type MockItemGetter struct {
	expectedBehavior ExpectedBehavior
}

func (mig *MockItemGetter) GetItems() []*domain.ItemRead {
	return testItems
}

func (mig *MockItemGetter) GetItemById(id string) (*domain.ItemRead, usecase.ItemGetterErrors) {
	switch mig.expectedBehavior {
	case Normal:
		return testItems[0], nil
	case NotFound:
		return nil, &ucustomerr.DataNotFoundError{Msg: "itemが見つかりませんでした。"}
	case TooManyResultsFound:
		return nil, &ucustomerr.TooManyResultsFoundError{Msg: "複数のitemが見つかりました。"}
	}

	return nil, nil
}

func TestItemController_GetItems(t *testing.T) {
	// arrange
	ic := controller.NewItemsController(&MockItemGetter{})

	// act
	actual, err := ic.GetItems(context.TODO(), api.GetItemsRequestObject{})
	if err != nil {
		t.Fatal(err)
	}
	// assert
	if result, ok := actual.(api.GetItems200JSONResponse); !ok {
		t.Fatalf("expected type: %T, but got %T", api.GetItems200JSONResponse{}, result)
	}

	expected := expectedItems

	for i, item := range actual.(api.GetItems200JSONResponse).Items {
		expected := expected[i]
		assertItem(t, expected, item)
	}
}

func TestItemController_GetItemsById_正常系(t *testing.T) {
	// arrange
	ic := controller.NewItemsController(&MockItemGetter{})

	// act
	actual, err := ic.GetItemsById(context.TODO(), api.GetItemsByIdRequestObject{ItemId: "1"})
	if err != nil {
		t.Fatal(err)
	}
	// assert
	if result, ok := actual.(api.GetItemsById200JSONResponse); !ok {
		t.Fatalf("expected type: %T, but got %T", api.GetItemsById200JSONResponse{}, result)
	}

	expected := expectedItems[0]
	assertItem(t, expected, actual.(api.GetItemsById200JSONResponse).Item)
}

func TestItemController_GetItemsById_対象のIdのitemがない(t *testing.T) {
	// arrange
	ic := controller.NewItemsController(&MockItemGetter{expectedBehavior: NotFound})

	// act
	actual, err := ic.GetItemsById(context.TODO(), api.GetItemsByIdRequestObject{ItemId: "3"})
	if err != nil {
		t.Fatal(err)
	}

	// assert
	if result, ok := actual.(api.GetItemsById404JSONResponse); !ok {
		t.Fatalf("expected type: %T, but got %T", api.GetItemsById404JSONResponse{}, result)
	}

	expected := "itemが見つかりませんでした。"
	if actual.(api.GetItemsById404JSONResponse).Message != expected {
		t.Errorf("expected: %s, but got %s", expected, actual.(api.GetItemsById404JSONResponse).Message)
	}
}

func TestItemController_GetItemsById_対象のIdのitemが複数ある(t *testing.T) {
	// arrange
	ic := controller.NewItemsController(&MockItemGetter{expectedBehavior: TooManyResultsFound})

	// act
	actual, err := ic.GetItemsById(context.TODO(), api.GetItemsByIdRequestObject{ItemId: "1"})
	if err != nil {
		t.Fatal(err)
	}

	// assert
	if result, ok := actual.(api.GetItemsById500JSONResponse); !ok {
		t.Fatalf("expected type: %T, but got %T", api.GetItemsById500JSONResponse{}, result)
	}

	expected := "複数のitemが見つかりました。"
	if actual.(api.GetItemsById500JSONResponse).Message != expected {
		t.Errorf("expected: %s, but got %s", expected, actual.(api.GetItemsById500JSONResponse).Message)
	}
}

func assertItem(t *testing.T, expected api.Item, actual api.Item) {
	t.Helper()
	if actual.Id != expected.Id {
		t.Errorf("expected: %s, but got %s", expected.Id, actual.Id)
	}
	if actual.ItemName != expected.ItemName {
		t.Errorf("expected: %s, but got %s", expected.ItemName, actual.ItemName)
	}
	if actual.JanCode != expected.JanCode {
		t.Errorf("expected: %s, but got %s", expected.JanCode, actual.JanCode)
	}
	if actual.Price != expected.Price {
		t.Errorf("expected: %d, but got %d", expected.Price, actual.Price)
	}
	if actual.CategoryId != expected.CategoryId {
		t.Errorf("expected: %d, but got %d", expected.CategoryId, actual.CategoryId)
	}
	if actual.SeriesId != expected.SeriesId {
		t.Errorf("expected: %d, but got %d", expected.SeriesId, actual.SeriesId)
	}
	if actual.Stock != expected.Stock {
		t.Errorf("expected: %d, but got %d", expected.Stock, actual.Stock)
	}
	if actual.Discontinued != expected.Discontinued {
		t.Errorf("expected: %t, but got %t", expected.Discontinued, actual.Discontinued)
	}
	if actual.ReleasedDate.Time != expected.ReleasedDate.Time {
		t.Errorf("expected: %s, but got %s", expected.ReleasedDate.Time, actual.ReleasedDate.Time)
	}
}
