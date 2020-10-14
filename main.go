package main

import (
	"linq-test/model"
)

func main() {
	Categories := []model.Category{
		{1, "gin"},
		{2, "vodka"},
		{3, "whisky"},
		{4, "brandy"},
		{5, "rum"},
	}

	Accounts := []model.Account{
		{ID: 1, Name: "Ray"},
		{ID: 2, Name: "Amy"},
	}

	Products := []model.Product{
		{ID: 1, CreatorID: 1, Name: "gin tonic", Category: Categories[0], Currency: model.Currency{Value: 250, Type: 1}},
		{ID: 2, CreatorID: 1, Name: "gin fizz", Category: Categories[0], Currency: model.Currency{Value: 280, Type: 1}},
		{ID: 3, CreatorID: 1, Name: "negroni", Category: Categories[0], Currency: model.Currency{Value: 290, Type: 1}},
		{ID: 4, CreatorID: 1, Name: "gin martin", Category: Categories[0], Currency: model.Currency{Value: 350, Type: 1}},
		{ID: 5, CreatorID: 2, Name: "vesper martin", Category: Categories[1], Currency: model.Currency{Value: 350, Type: 1}},
		{ID: 6, CreatorID: 2, Name: "vodka lime", Category: Categories[1], Currency: model.Currency{Value: 280, Type: 1}},
		{ID: 7, CreatorID: 2, Name: "cosmopolitan", Category: Categories[1], Currency: model.Currency{Value: 300, Type: 1}},
		{ID: 8, CreatorID: 3, Name: "sidecar", Category: Categories[3], Currency: model.Currency{Value: 350, Type: 1}},
	}

	ProductInnerJoinCreatorInfo(Products, Accounts)
	ProductLeftJoinCreatorInfo(Products, Accounts)

	SearchByCategory(Products)
	SearchByNameLike(Products)
	SearchByNameLikeAndCategory(Products)
	SearchByCategoryAndPrice(Products)
}
