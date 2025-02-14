package prices

import (
	"fmt"
	"PriceCalculator/conversion"
	"PriceCalculator/FileManager"
)


type TaxIncludedPriceJob struct {
	TaxRate          float64
	InputPrices      []float64
	TaxIncludedPrice map[string]float64
}

func (job  *TaxIncludedPriceJob) LoadPrices (){
	line,err:=filemanager.ReadLines("prices.txt")
	prices,err:=conversion.StringToFloats(line)
	if err!=nil{
		fmt.Println("Error while reading file")
		fmt.Println(err) 
		return
	}
job.InputPrices=prices
}

func (job *TaxIncludedPriceJob) Process() {
	job.LoadPrices()
	result := make(map[string]string)
	for _, price := range job.InputPrices {
		taxinclprice:=price * (1 + job.TaxRate)
		result[fmt.Sprintf("%.2f",price)] = fmt.Sprintf("%.2f",taxinclprice)
	}
	fmt.Println(result);
}

// Constructor generally the naming convention uses "new" keyword
func NewTaxIncludedPriceJob(taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		InputPrices: []float64{10, 20, 20},
		TaxRate:taxRate,
	}

}