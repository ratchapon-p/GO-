package main

import "fmt"

func main() {
	userNames := make([]string, 2)

	userNames[0] = "Julie"
	userNames[1] = "Charlie"

	userNames = append(userNames, "Max")
	userNames = append(userNames, "Manuel")

	fmt.Println(userNames)

}