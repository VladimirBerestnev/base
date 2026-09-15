package main

import (
	"errors"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	for {
		operation, err1 := chooseOperation()
		if err1 != nil {
			fmt.Println(err1)
			continue
		}
		if operation == 4 {
			break
		}
		numbers := enterNumbers()
		switch operation {
		case 1:
			AVG(numbers)
		case 2:
			SUM(numbers)
		case 3:
			MED(numbers)
		default:
			fmt.Println("Что-то пошло не так")
		}
		fmt.Println("Нужен еще расчет? Y/n")
		change := ""
		fmt.Scan(&change)
		if change != "Y" && change != "y" {
			break
		}
	}
}

func chooseOperation() (int, error) {
	operation := 0
	fmt.Println("Выберите операцию с числами: ")
	fmt.Println("1 - AVG(среднее) ")
	fmt.Println("2 - SUM(сумма)")
	fmt.Println("3 - MED(медиана)")
	fmt.Println("4 - Выход")
	fmt.Scan(&operation)
	if operation != 1 && operation != 2 && operation != 3 && operation != 4 {
		return 0, errors.New("Неверный ввод операции. Повторите ввод")
	}
	return operation, nil
}

func enterNumbers() []int {
	numbersString := ""
	numbersInt := []int{}
	for {
		fmt.Println("Введите числа через запятую: ")
		fmt.Scan(&numbersString)
		numbersSplit := strings.Split(numbersString, ",")
		for _, num := range numbersSplit {
			numTrim := strings.TrimSpace(num)
			number, err := strconv.Atoi(numTrim)
			if err != nil {
				fmt.Println("Не могу преобразовать в число ", numTrim)
				fmt.Println("Повторите ввод")
				continue
			}
			numbersInt = append(numbersInt, number)
		}
		return numbersInt
	}
}

func AVG(numbers []int) {
	sum := 0
	var result float64
	for _, number := range numbers {
		sum += number
	}
	result = float64(sum) / float64(len(numbers))
	fmt.Println("Среднее значение введенных чисел: ", result)
}
func SUM(numbers []int) {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	fmt.Println("Сумма введенных чисел: ", sum)
}
func MED(numbers []int) {
	var MED float64
	sort.Ints(numbers)
	if len(numbers)%2 == 0 {
		MED = (float64(numbers[len(numbers)/2-1]) + float64(numbers[len(numbers)/2])) / 2
	} else {
		MED = float64(numbers[len(numbers)/2])
	}
	fmt.Println(MED)
}
