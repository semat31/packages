// Package store provides types and methods
// commonly required for online sales
package store

// standardTax standerd tax and threshold
var standardTax = newTaxRate(0.25, 20)

// Product describes an item for sale
type Product struct {
	Name     string
	Category string
	ptice    float64
}

// NewProduct Product structure constructor
func NewProduct(name, category string, price float64) *Product {
	return &Product{name, category, price}
}

// Price ...
func (p *Product) Price() float64 {
	return standardTax.calcTax(p)
}

// SetPrice ...
func (p *Product) SetPrice(newPrice float64) {
	p.ptice = newPrice
}
