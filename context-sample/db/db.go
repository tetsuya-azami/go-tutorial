package db

import (
	"context"
	"fmt"
)

type MyDB struct {
}

type SearchResult struct {
	Data string
	Err  error
}

var DefaultDB MyDB = MyDB{}

func (db *MyDB) Search(ctx context.Context, userID int) <-chan SearchResult {
	result := make(chan SearchResult)
	go func(ctx context.Context) {
		result <- SearchResult{
			Data: fmt.Sprintf("User ID: %d", userID),
			Err:  nil,
		}
	}(ctx)

	return result
}
