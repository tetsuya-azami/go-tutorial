package domain

import "time"

type ItemRead struct {
	id           string
	janCode      string
	itemName     string
	price        int
	categoryId   int
	seriesId     int
	stock        int
	discontinued bool
	releaseDate  time.Time
	deletedAt    time.Time
}

func NewItemRead(
	id string,
	janCode string,
	itemName string,
	price int,
	categoryId int,
	seriesId int,
	stock int,
	discontinued bool,
	releaseDate time.Time,
	deletedAt time.Time,
) (*ItemRead, error) {
	return &ItemRead{
		id:           id,
		janCode:      janCode,
		itemName:     itemName,
		price:        price,
		categoryId:   categoryId,
		seriesId:     seriesId,
		stock:        stock,
		discontinued: discontinued,
		releaseDate:  releaseDate,
		deletedAt:    deletedAt,
	}, nil
}

func (itemRead *ItemRead) IsDeleted() bool {
	return !itemRead.deletedAt.IsZero()
}
