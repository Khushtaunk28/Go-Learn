package prices

import (
	"PriceCalculator/conversion"
	"PriceCalculator/iomanager"
	"fmt"
)


type TaxIncludedPriceJob struct {
	IOManager iomanager.IOManager `json:"-"`
	TaxRate          float64 `"json:"tax_rate`
	InputPrices      []float64 `json:"input_prices"`
	TaxIncludedPrice map[string]float64 `json:"tax_included_Price"`
}

func (job  *TaxIncludedPriceJob) LoadPrices (){
	line,err:=job.IOManager.ReadLines()
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
	//to store with dynamic names
	job.IOManager.WriteResult(result)
	//fmt.Println(result);
}

// Constructor generally the naming convention uses "new" keyword
func NewTaxIncludedPriceJob(iom iomanager.IOManager,taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		IOManager: iom ,
		InputPrices: []float64{10, 20, 20},
		TaxRate:taxRate,
	}
}