package store

const defaultTaxRate float64 = 0.2
const minThreshold = 10

// categoryMaxPrice ...
var categoryMaxPrice = map[string]float64{
	"Watersports": 250,
	"Soccer":      150,
	"Chess":       50,
}

//
func init() {
	for category, price := range categoryMaxPrice {
		categoryMaxPrice[category] = price + (price * defaultTaxRate)
	}
}

// taxRate ...
type taxRate struct {
	rate      float64
	threshold float64
}

// newTaxRate: taxRate structure construcctor.
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
func (taxRate *taxRate) calcTax(product *Product) (price float64) {
	if product.ptice > taxRate.threshold {
		price = product.ptice + (product.ptice * taxRate.rate)
	} else {
		price = product.ptice
	}

	if max, ok := categoryMaxPrice[product.Category]; ok && price > max {
		price = max
	}
	return
}
