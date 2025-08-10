package main

import (
	"errors"
	"fmt"
	"strings"
)

const (
	EURtoUSD = 1.08
	USDtoEUR = 0.93
	EURtoRUB = 100.0
	RUBtoEUR = 0.01
	USDtoRUB = 92.0
	RUBtoUSD = 0.011
)

func main() {
	for {
		inputCurrency, amount, outputCurrency := clientInput()
		result := calculateCurrency(strings.ToUpper(inputCurrency), amount, strings.ToUpper(outputCurrency))
		fmt.Printf("Результат: %.2f %s\n\n", result, strings.ToUpper(outputCurrency))
		if !checkRepeatCalculate() {
			break
		}
	}
}

func clientInput() (string, float64, string) {
	var amount float64
	var inputCurrency, outputCurrency string
	fmt.Println("___ Калькулятор валют ___")
	fmt.Println("Доступные валюты: EUR, USD, RUB")
	for {
		fmt.Println("Введите исходную валюту EUR/USD/RUB")
		_, _ = fmt.Scan(&inputCurrency)
		if _, err := checkCurrency(inputCurrency); err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println("Укажите количество купюр: ")
		_, _ = fmt.Scan(&amount)
		if _, err := checkAmount(amount); err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println("Введите целувую валюту EUR/USD/RUB")
		_, _ = fmt.Scan(&outputCurrency)
		if _, err := checkCurrency(outputCurrency); err != nil {
			fmt.Println(err)
			continue
		}
		break
	}
	return inputCurrency, amount, outputCurrency

}

func calculateCurrency(inputCurrency string, amount float64, outputCurrency string) float64 {
	switch {
	case inputCurrency == "EUR" && outputCurrency == "USD":
		return amount * EURtoUSD
	case inputCurrency == "USD" && outputCurrency == "EUR":
		return amount * USDtoEUR
	case inputCurrency == "EUR" && outputCurrency == "RUB":
		return amount * EURtoRUB
	case inputCurrency == "RUB" && outputCurrency == "EUR":
		return amount * RUBtoEUR
	case inputCurrency == "USD" && outputCurrency == "RUB":
		return amount * USDtoRUB
	case inputCurrency == "RUB" && outputCurrency == "USD":
		return amount * RUBtoUSD
	case inputCurrency == outputCurrency:
		return amount
	default:
		return amount
	}
}

func checkAmount(amount float64) (float64, error) {
	if amount <= 0 {
		return 0, errors.New("Неверное значение! Сумма должна быть больше 0")
	}
	return amount, nil
}

func checkCurrency(currency string) (string, error) {
	currency = strings.ToUpper(strings.TrimSpace(currency))
	if currency != "EUR" && currency != "RUB" && currency != "USD" {
		return currency, errors.New("Неверная валюта. Доступные варианты: EUR, USD, RUB")
	}
	return currency, nil
}

func checkRepeatCalculate() bool {
	var userAnswer string
	fmt.Println("Произвести повторный расчет?(y/n)")
	_, _ = fmt.Scan(&userAnswer)
	if userAnswer == "y" || userAnswer == "Y" {
		return true
	}
	return false
}
