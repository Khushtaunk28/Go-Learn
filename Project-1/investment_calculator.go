package main

import (
	"fmt"
	"math"
)

func main() {
	//cont wont change value
	const inflationRate = 2.5
	//same line declaration but of same type
	var invAmount float64
	//shorcut to assignvar rate = 5.5
	var rate float64 = 5.5
	years := 10.0

	//take input from user,we use pointers
	fmt.Print("Enter the Investment Amount : ")
	fmt.Scan(&invAmount)

	futvalue := invAmount * math.Pow((1+rate/100), float64(years))
	futRealValue := futvalue / math.Pow(1+inflationRate/100, years)
	fmt.Println(futvalue)
	fmt.Println(futRealValue)
	fmt.Printf(`Future Value is=\n
	%.2f`, futvalue)
}
