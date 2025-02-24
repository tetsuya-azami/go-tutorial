package data

import "time"

type ItemData struct {
	Id           string
	JanCode      string
	ItemName     string
	Price        int
	CategoryId   int
	SeriesId     int
	Stock        int
	Discontinued bool
	ReleaseDate  time.Time
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    time.Time
}
