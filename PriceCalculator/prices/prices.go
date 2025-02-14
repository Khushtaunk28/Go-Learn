package prices

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type TaxIncludedPriceJob struct {
	TaxRate          float64
	InputPrices      []float64
	TaxIncludedPrice map[string]float64
}

func (job  *TaxIncludedPriceJob) LoadPrices (){
	file,err:= os.Open("prices.txt")
	if(err!=nil){
		fmt.Println("Error while opening file")
		fmt.Println(err)
		return
	}
	scanner:=bufio.NewScanner(file)
	var line []string
	for scanner.Scan(){
		line =append(line,scanner.Text())

	}
	err=scanner.Err()
	if err!=nil{
		fmt.Println("Error while reading file")
		fmt.Println(err)
		file.Close()
		return
	}
	prices :=make([]float64,len(line))
	for lineIndex,line :=range line{
		floatPrice,err:=strconv.ParseFloat(line,64)
	
	if err!=nil{
		fmt.Println("Error while converting file")
		fmt.Println(err)
		file.Close()
		return
	}
	prices[lineIndex]=floatPrice
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
		TaxRate:     taxRate,
	}

}