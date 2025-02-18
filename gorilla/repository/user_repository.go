package repository

import "gorilla-tutorial/domain"

var Users = map[int]*domain.User{
	1: {
		ID:      1,
		Name:    "John",
		Country: "USA",
	},
	2: {
		ID:      2,
		Name:    "Jane",
		Country: "Canada",
	},
	3: {
		ID:      2,
		Name:    "Jane",
		Country: "Canada",
	},
}
