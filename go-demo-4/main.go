package main

import (
	"bufio"
	"fmt"
	"go-demo-4/password/account"
	"os"

	"github.com/fatih/color"
)

func main() {
	fmt.Println("___Менеджер Аккаунтов___")
	fmt.Println("___Выберте пункт меню___")
	fmt.Println(`1 - Создать Аккаунт
2 - Найти Аккаунт
3 - Удалить Аккаунт
4 - Выход из Меню
`)
	vault := account.NewVault()
Menu:
	for {
		input := clientInput()
		switch input {
		case "1":
			createAccount(vault)
		case "2":
			findAccount(vault)
		case "3":
			removeAccount(vault)
		default:
			break Menu
		}
	}
}
func clientInput() string {
	fmt.Println("Введите пункт меню:")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()
	return input
}

func createAccount(vault *account.Vault) {
	login := promptData("Input Login")
	password := promptData("Input Password")
	urlString := promptData("Input URL")
	myAccount, err := account.NewAccount(login, password, urlString)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	vault.AddAccount(*myAccount)
}

func findAccount(vault *account.Vault) {
	url := promptData("Input URL")
	accounts := vault.FindAccountsByURL(url)
	if len(accounts) == 0 {
		color.Red("Account not found")
	}
	for _, account := range accounts {
		account.Output()
	}

}

func removeAccount(vault *account.Vault) {
	url := promptData("Input URL")
	isDeleted := vault.DeleteAccountByUrl(url)
	if isDeleted {
		color.Green("Account deleted")
	} else {
		color.Red("Account not deleted")
	}
}

func promptData(prompt string) string {
	fmt.Println(prompt + ":")
	var result string
	_, _ = fmt.Scanln(&result)
	return result
}
