package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

const USDtoEUR = 0.95
const EURtoUSD = 1.05
const USDtoRUB = 75
const EURtoRUB = 80
const RUBtoEUR = 0.0125
const RUBtoUSD = 0.01333

var currency = map[string]map[string]float64{
	"EUR": {"USD": EURtoUSD, "RUB": EURtoRUB},
	"USD": {"EUR": USDtoEUR, "RUB": USDtoRUB},
	"RUB": {"USD": RUBtoUSD, "EUR": RUBtoEUR},
}

func main() {
	userInput()
}

func userInput() {
	for {
		userCurrency := getUserCurrency()
		needCurrency := getNeedCurrency(userCurrency)
		money, err := getMoney()
		if err != nil {
			fmt.Println(err)
			getMoney()
		}
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

	value := float64(money) * currency[userCurrency][needCurrency]

	fmt.Printf("В результате конвертации %d %s Вы получите: %.2f %s\n", money, userCurrency, value, needCurrency)
}

func getUserCurrency() string {
	userCurrency := ""
	fmt.Println("Введите Вашу валюту")
	fmt.Println("RUB - Рубли, USD - Доллары")
	fmt.Println("EUR - Евро, EXIT - Выход")
	fmt.Scan(&userCurrency)
	userCurrency = strings.ToUpper(userCurrency)
	if userCurrency != "EXIT" && userCurrency != "RUB" && userCurrency != "EUR" && userCurrency != "USD" {
		fmt.Println("Неверный ввод данных")
		getUserCurrency()
	}
	return userCurrency
}
func getNeedCurrency(userCurrency string) string {
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
	default:
		fmt.Println("Неверный ввод. Введите заново")
	}
	fmt.Scan(&needCurrency)
	needCurrency = strings.ToUpper(needCurrency)
	if needCurrency != "EXIT" && needCurrency != "RUB" && needCurrency != "EUR" && needCurrency != "USD" {
		fmt.Println("Неверный ввод данных")
		getNeedCurrency(userCurrency)
	}

	return needCurrency
}

func getMoney() (int, error) {
	money := ""
	fmt.Println("Введите количество денег: ")
	fmt.Scan(&money)
	userMoney, err := strconv.Atoi(money)
	if err != nil {
		return 0, errors.New("Неверный ввод данных")
	}
	return userMoney, nil
}
