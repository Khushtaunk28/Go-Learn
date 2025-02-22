package main

import (
	//"PriceCalculator/cmdmanager"
	filemanager "PriceCalculator/fileManager"
	"PriceCalculator/prices"
	"fmt"
)

func main() {
	taxRates := []float64{0, 0.05, 0.1, 0.15}
	for _, taxRates := range taxRates {
		fm:=filemanager.NewFileManager("prices.txt",fmt.Sprintf("result_%.2f.json",taxRates*100))
		//cmdm:=cmdmanager.New()
		priceJob:=prices.NewTaxIncludedPriceJob(fm,taxRates)
		err:=priceJob.Process()
		if err!=nil{
			fmt.Println("Job not proccessed")
			fmt.Println(err)
		}
	}
}