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

// 全プロパティのgetterを生成してください
func (itemRead *ItemRead) Id() string {
	return itemRead.id
}

func (itemRead *ItemRead) JanCode() string {
	return itemRead.janCode
}

func (itemRead *ItemRead) ItemName() string {
	return itemRead.itemName
}

func (itemRead *ItemRead) Price() int {
	return itemRead.price
}

func (itemRead *ItemRead) CategoryId() int {
	return itemRead.categoryId
}

func (itemRead *ItemRead) SeriesId() int {
	return itemRead.seriesId
}

func (itemRead *ItemRead) Stock() int {
	return itemRead.stock
}

func (itemRead *ItemRead) Discontinued() bool {
	return itemRead.discontinued
}

func (itemRead *ItemRead) ReleaseDate() time.Time {
	return itemRead.releaseDate
}

func (itemRead *ItemRead) DeletedAt() time.Time {
	return itemRead.deletedAt
}
