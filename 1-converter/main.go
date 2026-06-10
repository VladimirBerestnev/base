package main

import "fmt"

func main() {
	const USDtoEUR = 0.95
	const USDtoRUB = 75
	EURtoRUB := USDtoRUB / USDtoEUR
	fmt.Println(EURtoRUB)
}
