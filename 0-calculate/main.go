package main

import (
	"errors"
	"fmt"
	"math"
)

func main() {
	fmt.Println("___ Калькулятор индекса массы тела ___")
	for {
		userHeight, userKg := getUserInput()
		IMT, err := calcucalteIMT(userKg, userHeight)
		if err != nil {
			fmt.Println(err)
			continue
		}
		outputResult(IMT)
		repeatCalculation := checkRepeatCalc()
		if !repeatCalculation {
			break
		}
	}
}

func outputResult(imt float64) {
	result := fmt.Sprintf("Ваш индекс массы тела: %.0f\n", imt)
	fmt.Print(result)
	switch {
	case imt < 16:
		fmt.Println("У вас сильный дефицит массы тела")
	case imt < 18.5:
		fmt.Println("У вас дефицит массы тела")
	case imt < 25:
		fmt.Println("У вас нормальный вес")
	case imt < 30:
		fmt.Println("У вас избыточный вес")
	default:
		fmt.Println("У вас степерь ожирения")
	}
}

func calcucalteIMT(userKg, userHeight float64) (float64, error) {
	const IMTPower = 2
	if userHeight <= 0 || userKg <= 0 {
		return 0, errors.New("WRONG_INPUTS_VALUE")
	}
	IMT := userKg / math.Pow((userHeight/100), IMTPower)
	return IMT, nil
}

func getUserInput() (float64, float64) {
	var userKg float64
	var userHeight float64
	fmt.Print("Введите свой рост в сантиметрах: ")
	fmt.Scan(&userHeight)
	fmt.Print("Введите свой вес в килограммах: ")
	fmt.Scan(&userKg)
	return userHeight, userKg
}

func checkRepeatCalc() bool {
	var userAnswer string
	fmt.Println("Произвести повторный расчет?(y/n)")
	_, _ = fmt.Scan(&userAnswer)
	if userAnswer == "y" || userAnswer == "Y" {
		return true
	}
	return false
}
