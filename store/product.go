// Package store provides types and methods
// commonly required for online sales
package store

// Product describes an item for sale
type Product struct {
	Name     string
	Category string
	ptice    float64
}

// NewProduct ...
func NewProduct(name, category string, price float64) *Product {
	return &Product{name, category, price}
}

// Price ...
func (p *Product) Price() float64 {
	return p.ptice
}

// SetPrice ...
func (p *Product) SetPrice(newPrice float64) {
	p.ptice = newPrice
}
