package main

import (
	"fmt"
	"errors"
)

func main()  {
	revenue, err := getUserInput("Revenue : ")

	if err != nil{
		fmt.Println(err)
		return
		// panic(err)
	}

	expenses, err := getUserInput("Expense : ")

	if err != nil{
		fmt.Println(err)
		return

		// panic(err)

	}

	taxRate, err := getUserInput("Tax Rate : ")

	if err != nil{
		fmt.Println(err)
		return

		// panic(err)

	}

    ebt, profit, ratio := calculateFinancials(revenue, expenses, taxRate)


	fmt.Printf("%.1f \n", ebt)
	fmt.Printf("%.1f \n", profit)
	fmt.Printf("%.1f \n", ratio)

}

func calculateFinancials(revenue, expense, taxRate float64) (float64, float64, float64){
	ebt := revenue - expense
	profit := ebt * (1 - taxRate/100)
	ratio := ebt/profit
	return ebt, profit, ratio
}

func getUserInput(infoText string) (float64, error) {
	var userInput float64
	fmt.Print(infoText)
	fmt.Scan(&userInput)

	if userInput <= 0{
		return 0, errors.New("Value must be positive number.")
	}
	return userInput, nil

}

