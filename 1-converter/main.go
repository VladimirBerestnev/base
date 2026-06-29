package main

import (
	"fmt"
	"strings"
)

const USDtoEUR = 0.95
const EURtoUSD = 1.05
const USDtoRUB = 75
const RUBtoEUR = 80

func main() {

	EURtoRUB := USDtoRUB / USDtoEUR
	fmt.Println(EURtoRUB)

	userInput()
}

func userInput() {
	userCurrency := ""
cycle:
	for {
		fmt.Println("Введите Вашу валюту")
		fmt.Println("RUB - Рубли")
		fmt.Println("USD - Доллары")
		fmt.Println("EUR - Евро")
		fmt.Println("EXIT - Выход")
		fmt.Scan(&userCurrency)
		userCurrency = strings.ToUpper(userCurrency)
		needCurrency := ""
		fmt.Println("Введите валюту, в которую хотите перевести деньги: ")

		switch userCurrency {
		case "RUB":
			fmt.Println("EUR - Евро")
			fmt.Println("USD - Доллары")
		case "USD":
			fmt.Println("RUB - Рубли")
			fmt.Println("EUR - Евро")
		case "EUR":
			fmt.Println("RUB - Рубли")
			fmt.Println("USD - Доллары")
		case "EXIT":
			break cycle
		default:
			fmt.Println("Неверный ввод. Введите заново")
		}
		fmt.Scan(&needCurrency)
		needCurrency = strings.ToUpper(needCurrency)
		money := 0
		fmt.Println("Введите количество денег: ")
		fmt.Scan(&money)
		exchange(money, userCurrency, needCurrency)

		fmt.Println("Нужен еще расчет? Y/n")
		change := ""
		fmt.Scan(&change)
		if change == "n" || change == "N" {
			break
		}
	}
}

func exchange(money int, userCurrency string, needCurrency string) {
	var value float64
	if userCurrency == "USD" && needCurrency == "EUR" {
		value = float64(money) * USDtoEUR
	} else if userCurrency == "USD" && needCurrency == "RUB" {
		value = float64(money) * USDtoRUB
	} else if userCurrency == "RUB" && needCurrency == "USD" {
		value = float64(money) / USDtoRUB
	} else if userCurrency == "RUB" && needCurrency == "EUR" {
		value = float64(money) * RUBtoEUR
	} else if userCurrency == "EUR" && needCurrency == "RUB" {
		value = float64(money) * USDtoRUB / USDtoEUR
	} else if userCurrency == "EUR" && needCurrency == "USD" {
		value = float64(money) * EURtoUSD
	} else {
		fmt.Println("Неверные данные для расчета")
	}
	fmt.Printf("В результате конвертации %d %s Вы получите: %.2f %s\n", money, userCurrency, value, needCurrency)
}
