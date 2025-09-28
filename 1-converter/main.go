package main

import (
	"errors"
	"fmt"
	"strings"
)

type CurrencyRates = map[string]float64

var currencyMap = &CurrencyRates{
	"USD": 1.0,
	"EUR": 0.85,
	"RUB": 83.5,
}

func main() {
	for {
		inputCurrency, amount, outputCurrency := clientInput()
		result := calculateCurrency(strings.ToUpper(inputCurrency), amount, strings.ToUpper(outputCurrency), currencyMap)
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

func calculateCurrency(inputCurrency string, amount float64, outputCurrency string, m *CurrencyRates) float64 {
	return (amount / (*m)[inputCurrency]) * (*m)[outputCurrency]
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
