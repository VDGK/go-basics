// Декомпозиция задачи
// 1. Функция, которая будет принимать операцию (AVG,SUM,MED)
// 2. Функция, которая будет принимать числовые значения и наполнять слайс
// 3. Функция AVG
// 4. Функция SUM
// 5. Функция MED
// 6. Обработка ошибок и проверка вводных данных

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	method := methodType()
	transactions := scanTransactions()

	switch {
	case method == "SUM":
		sum := sumTransactions(transactions)
		fmt.Printf("Результат %.2f", sum)
	case method == "AVG":
		avg := avgTransactions(transactions)
		fmt.Printf("Результат %.2f", avg)
	case method == "MED":
		med := medTransactions(transactions)
		fmt.Printf("Результат %.2f", med)
	}
}

func methodType() string {
	var method string
	fmt.Println("Введите желаемую операцию (AVG/SUM/MED):")
	_, _ = fmt.Scan(&method)
	return strings.ToUpper(method)
}

func scanTransactions() []float64 {
	fmt.Println("Введите значения транзакций через запятую: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()
	numStr := strings.Split(input, ",")
	transactions := []float64{}
	for _, value := range numStr {
		num, err := strconv.ParseFloat(strings.TrimSpace(value), 64)
		if err != nil {
			fmt.Printf("Пропускаем нечисловое значение '%s'\n", value)
			continue
		}
		transactions = append(transactions, num)
	}
	return transactions
}

func sumTransactions(transactions []float64) float64 {
	var sum float64
	for _, value := range transactions {
		sum += value
	}
	return sum
}

func avgTransactions(transactions []float64) float64 {
	var total float64
	for _, value := range transactions {
		total += value
	}
	return total / float64(len(transactions))
}

func medTransactions(transactions []float64) float64 {
	sort.Float64s(transactions)

	medNumber := len(transactions) / 2

	if len(transactions)%2 == 0 {
		return transactions[medNumber]
	}

	return (transactions[medNumber] + transactions[medNumber-1]) / 2
}
