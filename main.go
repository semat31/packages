package main

import (
	_ "packages/data"
	currencyFmt "packages/fmt"
	"packages/store"
	"packages/store/cart"

	"github.com/fatih/color"
)

func main() {
	product := store.NewProduct("Kayak", "Watersports", 279)

	cart := cart.Cart{
		CustomerName: "Alice",
		Products:     []store.Product{*product},
	}

	color.Green("Name: " + cart.CustomerName)
	color.Cyan("Total: " + currencyFmt.ToCurrency(cart.GetTotal()))
}
