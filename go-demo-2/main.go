package main

import (
	"fmt"
)

func main() {
	transactions := []float64{}
	for {
		transaction := scanTransactions()
		if transaction == 0 {
			break
		}
		transactions = append(transactions, transaction)
	}
	sum := sumTransactions(transactions)
	fmt.Printf("Ваш баланс %.2f", sum)
}
func sumTransactions(transactions []float64) float64 {
	var sum float64
	for _, value := range transactions {
		sum += value
	}
	return sum
}
func scanTransactions() float64 {
	var transaction float64
	fmt.Println("Введите значение транзакции (n для выхода): ")
	_, _ = fmt.Scan(&transaction)
	return transaction
}
