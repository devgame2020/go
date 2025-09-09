package store

const defaultTaxRate float64 = 0.2
const minThreshold = 10

type taxRate struct {
	rate, threshold float64
}
