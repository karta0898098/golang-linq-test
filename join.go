package main

import (
	"github.com/ahmetb/go-linq/v3"
	"linq-test/model"
	"linq-test/util"
)

func ProductLeftJoinCreatorInfo(Products []model.Product, Accounts []model.Account) {
	productView := make([]model.ProductView, 0)
	linq.From(Products).GroupJoinT(
		linq.From(Accounts),
		func(p model.Product) int { return p.CreatorID },
		func(a model.Account) int { return a.ID },
		func(product interface{}, accounts []interface{}) interface{} {
			return linq.Group{
				Key:   product,
				Group: accounts,
			}
		},
	).SelectManyByT(
		func(x linq.Group) linq.Query {
			if len(x.Group) > 0 {
				return linq.From(x.Group)
			}
			return linq.From([]model.Account{{}})
		},
		func(acc interface{}, group interface{}) model.ProductView {
			p := group.(linq.Group).Key.(model.Product)
			a := acc.(model.Account)
			return model.ProductView{
				ID:       p.ID,
				Name:     p.Name,
				Creator:  a,
				Category: p.Category,
				Currency: p.Currency,
			}
		},
	).ToSlice(&productView)
	util.PrintViewResult("left join productView ==>", productView)
}

func ProductInnerJoinCreatorInfo(Products []model.Product, Accounts []model.Account) {
	productView := make([]model.ProductView, 0)
	linq.From(Products).JoinT(
		linq.From(Accounts),
		func(p model.Product) int {
			return p.CreatorID
		},
		func(a model.Account) int {
			return a.ID
		},
		func(x model.Product, y model.Account) model.ProductView {
			return model.ProductView{
				ID:       x.ID,
				Name:     x.Name,
				Creator:  y,
				Category: x.Category,
				Currency: x.Currency,
			}
		},
	).ToSlice(&productView)
	util.PrintViewResult("inner join productView ==>", productView)
}
