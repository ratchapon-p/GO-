package main

import (
	"fmt"
	"example.com/bank/fileops"

)

const accountBalanceFile = "balance.txt"

func main()  {

	var accountBalance, err = fileops.GetFloalFromFile(accountBalanceFile)
	if err != nil{
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("--------------------------------------------")
		panic("Can't continue sorry")
	}

	
	fmt.Println("Welcome to GO Bank!")

	for{

		presentOptions()
	
		var choice int
		fmt.Print("Your choice : ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Println("Your balance is : ",accountBalance)
			fileops.WriteFloatToFile(accountBalance, accountBalanceFile)
			break

		case 2 :
			fmt.Print("Your deposit : ")
			var depositAmount float64
			fmt.Scan(&depositAmount)

			if depositAmount <= 0{
				fmt.Println("Invalid amount. Must be greater than 0")
				continue
			}
			accountBalance += depositAmount
			fmt.Println("Balance updated! New amount : ", accountBalance)
			fileops.WriteFloatToFile(accountBalance, accountBalanceFile)


		case 3 :
			fmt.Print("Withdraw : ")
			var withdrawAmount float64
			fmt.Scan(&withdrawAmount)

			if withdrawAmount <= 0{
				fmt.Println("Invalid amount. Must be greater than 0")
				continue
			}
			if withdrawAmount > accountBalance{
				fmt.Println("Invalid amount. You can't withdraw more than you have.")
				continue
			}
			accountBalance -= withdrawAmount
			fmt.Println("Balance updated! New amount : ", accountBalance)
			fileops.WriteFloatToFile(accountBalance, accountBalanceFile)

		default:
			fmt.Println("Goodbye")
			fmt.Println("Thanks for choosing our bank")
			return
		}
	

	}

}

