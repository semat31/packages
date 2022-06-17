package store

const defaultTaxRate float64 = 0.2
const minThreshold = 10

// taxRate ...
type taxRate struct {
	rate      float64
	threshold float64
}

// newTaxRate ...
func newTaxRate(rate, threshold float64) *taxRate {
	if rate == 0 {
		rate = defaultTaxRate
	}
	if threshold < minThreshold {
		threshold = minThreshold
	}
	return &taxRate{rate, threshold}
}

// calcTax ...
func (taxRate *taxRate) calcTax(product *Product) float64 {
	if product.ptice > taxRate.threshold {
		return product.ptice + (product.ptice * taxRate.rate)
	}
	return product.ptice
}
