package main

import "fmt"

func main()  {
	prices := []float64{10.99, 8.99}
	prices[1] = 9.99

	prices = append(prices, 5.99,12.99, 212.12, 43.77, 48.99)
	prices = prices[1:]
	fmt.Println(prices)

	discountPrices := []float64{101.99, 80.99, 20.59}
	prices = append(prices, discountPrices...)
	fmt.Println(prices)

}



// type Product struct {
// 	title string
// 	id    string
// 	price float64
// }

// func main() {
// 	var productNames [4]string= [4]string{"A Book"}
// 	prices := [4]float64{10.99, 9.99, 45.99, 20.0}
// 	productNames[2] = "A Carpet"

// 	featurePrices := prices[1:]
// 	highlightedPrices := featurePrices[:1]
// 	fmt.Println(highlightedPrices)
// }