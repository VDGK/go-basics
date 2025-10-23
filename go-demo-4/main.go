package main

import (
	"fmt"
	"go-demo-4/password/account"
	"go-demo-4/password/files"
	"os"
)

func main() {
	login := promptData("Input Login")
	password := promptData("Input Password")
	urlString := promptData("Input URL")
	myAccount, err := account.NewAccountWithTimeStamp(login, password, urlString)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	myAccount.OutputPassword()
	files.WriteFile()
	fmt.Println(myAccount)
}
func promptData(prompt string) string {
	fmt.Println(prompt + ":")
	var result string
	_, _ = fmt.Scanln(&result)
	return result
}
