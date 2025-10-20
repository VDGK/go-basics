package main

import (
	"errors"
	"fmt"
	"math/rand/v2"
	"net/url"
	"os"
)

type account struct {
	login    string
	password string
	url      string
}

func (acc *account) outputPassword() {
	fmt.Println(acc.login, acc.password, acc.url)
}

func (acc *account) generatePassword(length int) {
	res := make([]rune, length)
	for i := range res {
		res[i] = letterRunes[rand.IntN(len(letterRunes))]
	}
	acc.password = string(res)
}

func newAccount(login, password, urlString string) (*account, error) {
	if login == "" {
		return nil, errors.New("login required")
	}
	_, err := url.ParseRequestURI(urlString)
	if err != nil {
		return nil, fmt.Errorf("неверный URL-адрес %q: %w", urlString, err)
	}
	newAcc := &account{
		login:    login,
		url:      urlString,
		password: password,
	}
	if password == "" {
		newAcc.generatePassword(20)
	}
	return newAcc, nil
}

var letterRunes = []rune("qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM1234567890-*!@#$%^&*()_+-=")

func main() {
	login := promptData("Input Login")
	password := promptData("Input Password")
	urlString := promptData("Input URL")
	myAccount, err := newAccount(login, password, urlString)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	myAccount.outputPassword()
	fmt.Println(myAccount.login, myAccount.password, myAccount.url)
}
func promptData(prompt string) string {
	fmt.Println(prompt + ":")
	var result string
	_, _ = fmt.Scanln(&result)
	return result
}
