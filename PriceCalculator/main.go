package main

import (
	"PriceCalculator/cmdmanager"
	"PriceCalculator/prices"
)

func main() {
	taxRates := []float64{0, 0.05, 0.1, 0.15}
	for _, taxRates := range taxRates {
		//fm:=filemanager.NewFileManager("prices.txt",fmt.Sprintf("result_%.2f.json",taxRates*100))
		cmdm:=cmdmanager.New()
		priceJob:=prices.NewTaxIncludedPriceJob(cmdm,taxRates)
		priceJob.Process()
	}
}