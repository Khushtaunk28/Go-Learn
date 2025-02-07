package main
import "fmt"

func main(){
	var productNames [4]string =[4]string {"hello"}
	prices:= [4]float64{10.1,2.3,5.333,765.00}

	productNames[2]="heyyy"
	
	fmt.Println(prices);
	fmt.Println(productNames);
	fmt.Println(prices[2])
//slice = to get the subset of an array (exclude the last index)
//everytime we use slice we just editing the original array, so any changes made in the sliced
//array would lead to changes in the original array
	featurePrices:= prices[1:3]
	fmt.Println(featurePrices)
	//cap=is the actually capacity of the array
	fmt.Println(len(featurePrices),cap(featurePrices))

}