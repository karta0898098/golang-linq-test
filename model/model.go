package model

import "strings"

// Category ...
type Category struct {
	ID   int
	Name string
}

// Currency ...
type Currency struct {
	Value int
	Type  int
}

// Account ...
type Account struct {
	ID   int
	Name string
}

// Product ...
type Product struct {
	ID        int
	Name      string
	CreatorID int
	Category  Category
	Currency  Currency
}

// ProductView ...
type ProductView struct {
	ID   int
	Name string

	Creator  Account
	Category Category
	Currency Currency
}

type QueryProductOption struct {
	Name     string
	Category string
	Price    *int
}

func (option QueryProductOption) Predicate() func(product Product) bool {
	return func(product Product) bool {
		return option.searchByCategory(product) &&
			option.searchByProductName(product) &&
			option.searchByGreaterThanEqualPrice(product)
	}
}

func (option QueryProductOption) searchByProductName(product Product) bool {
	if option.Name != "" {
		return strings.Contains(product.Name, option.Name)
	}
	return true
}

func (option QueryProductOption) searchByCategory(product Product) bool {
	if option.Category != "" {
		return product.Category.Name == option.Category
	}
	return true
}

func (option QueryProductOption) searchByGreaterThanEqualPrice(product Product) bool {
	if option.Price != nil {
		return product.Currency.Value >= *option.Price
	}
	return true
}


