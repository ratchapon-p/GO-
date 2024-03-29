package main

import (
	"fmt"
	"math"
)


func main()  {
	const inflationRate = 2.5
	var investmentAmount float64
	expectedReturnRate := 5.5
	var years float64

	fmt.Print("Investment Amount : ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Years : ")
	fmt.Scan(&years)

	fmt.Print("Except Return Rate : ")
	fmt.Scan(&expectedReturnRate)



	futureValue := investmentAmount * math.Pow(1 + expectedReturnRate / 100, years)
	futureRealValue := futureValue / math.Pow(1 + inflationRate/100, years)
	
	fmt.Println(futureValue)
	fmt.Println(futureRealValue)
}

