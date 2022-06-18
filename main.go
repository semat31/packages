package main

import (
	"fmt"
	currencyFmt "packages/fmt"
	"packages/store"
	"packages/store/cart"
)

func main() {
	product := store.NewProduct("Kayak", "Watersports", 279)
	product1 := store.NewProduct("Padle", "Watersports", 200)

	cart := cart.Cart{
		CustomerName: "Alice",
		Products:     []store.Product{*product, *product1},
	}

	fmt.Println("Name:", product.Name)
	fmt.Println("Category:", product.Category)
	fmt.Println("Price:", currencyFmt.ToCurrency(product.Price()))

	fmt.Println("Name:", cart.CustomerName)
	fmt.Println("Total:", currencyFmt.ToCurrency(cart.GetTotal()))
}
