package main

import (
	"PriceCalculator/prices"
)

func main() {
	taxRates := []float64{0, 0.05, 0.1, 0.15}
	for _, taxRates := range taxRates {
		priceJob:=prices.NewTaxIncludedPriceJob(taxRates)
		priceJob.Process()
	}
}