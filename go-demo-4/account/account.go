package account

import (
	"errors"
	"fmt"
	"math/rand/v2"
	"net/url"
	"time"

	"github.com/fatih/color"
)

type Account struct {
	Login     string    `json:"login"`
	Password  string    `json:"password"`
	Url       string    `json:"url"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (acc *Account) Output() {
	color.New(color.FgBlue).Println(acc.Login, acc.Password, acc.Url)
}

func (acc *Account) generatePassword(length int) {
	res := make([]rune, length)
	for i := range res {
		res[i] = letterRunes[rand.IntN(len(letterRunes))]
	}
	acc.Password = string(res)
}

func NewAccount(login, password, urlString string) (*Account, error) {
	if login == "" {
		return nil, errors.New("login required")
	}
	_, err := url.ParseRequestURI(urlString)
	if err != nil {
		return nil, fmt.Errorf("неверный URL-адрес %q: %w", urlString, err)
	}
	newAcc := &Account{
		Login:     login,
		Password:  password,
		Url:       urlString,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	if password == "" {
		newAcc.generatePassword(20)
	}
	return newAcc, nil
}

var letterRunes = []rune("qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM1234567890-*!@#$%^&*()_+-=")
