package main

import (
	"github.com/ahmetb/go-linq/v3"
	"linq-test/model"
	"linq-test/util"
)

func SearchByNameLikeAndCategory(Products []model.Product) {
	query := model.QueryProductOption{
		Name:     "martin",
		Category: "gin",
	}

	products := make([]model.Product, 0)
	linq.From(Products).
		WhereT(query.Predicate()).
		ToSlice(&products)

	util.PrintResult("Name like martin and Category is gin", products)
}

func SearchByCategory(Products []model.Product) {
	query := model.QueryProductOption{
		Category: "gin",
	}

	products := make([]model.Product, 0)
	linq.From(Products).
		WhereT(query.Predicate()).
		ToSlice(&products)

	util.PrintResult("Category is gin", products)
}

func SearchByNameLike(Products []model.Product) {
	query := model.QueryProductOption{
		Name: "martin",
	}

	products := make([]model.Product, 0)
	linq.From(Products).
		WhereT(query.Predicate()).
		ToSlice(&products)

	util.PrintResult("Name Like martin", products)
}

func SearchByCategoryAndPrice(Products []model.Product) {
	price := 280
	query := model.QueryProductOption{
		Category: "gin",
		Price:    &price,
	}

	products := make([]model.Product, 0)
	linq.From(Products).
		WhereT(query.Predicate()).
		ToSlice(&products)

	util.PrintResult("Category is gin And Price greater than equal 280", products)
}
