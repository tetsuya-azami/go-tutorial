package db

import (
	"context"
	"fmt"
	"time"
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
		defer close(result)
		time.Sleep(3 * time.Second)
		select {
		case <-ctx.Done():
			result <- SearchResult{
				Data: "",
				Err:  fmt.Errorf("context timeout"),
			}
		default:
			result <- SearchResult{
				Data: fmt.Sprintf("User ID: %d", userID),
				Err:  nil,
			}
		}
	}(ctx)

	return result
}
