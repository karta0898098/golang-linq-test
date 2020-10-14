package util

import (
	"fmt"
	"linq-test/model"
)

func PrintResult(result string, products []model.Product) {
	fmt.Printf(result + "\n")
	for i := 0; i < len(products); i++ {
		fmt.Printf("%#v\n", products[i])
	}
	fmt.Printf("\n")
}

func PrintViewResult(result string, products []model.ProductView) {
	fmt.Printf(result + "\n")
	for i := 0; i < len(products); i++ {
		fmt.Printf("%#v\n", products[i])
	}
	fmt.Printf("\n")
}
