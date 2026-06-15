package main

import "fmt"

func main() {
	const USDtoEUR = 0.95
	const USDtoRUB = 75
	EURtoRUB := USDtoRUB / USDtoEUR
	fmt.Println(EURtoRUB)

	money := userInput()
}

func userInput() int {
	var input int
	fmt.Scan(&input)
	return input
}

func exchange(money int, userCurrency string, needCurrency string) {

}
