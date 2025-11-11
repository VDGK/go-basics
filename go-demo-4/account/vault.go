package account

import (
	"encoding/json"
	"go-demo-4/password/files"
	"strings"
	"time"

	"github.com/fatih/color"
)

type Vault struct {
	Accounts  []Account `json:"accounts"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewVault() *Vault {
	file, err := files.ReadFromFile("data.json")
	if err != nil {
		return &Vault{
			Accounts:  []Account{},
			UpdatedAt: time.Now(),
		}
	}
	var vault Vault
	err = json.Unmarshal(file, &vault)
	if err != nil {
		color.Red(err.Error())
	}
	return &vault
}

func (vault *Vault) AddAccount(acc Account) {
	vault.Accounts = append(vault.Accounts, acc)
	vault.save()
}

func (vault *Vault) ToBytes() (data []byte, err error) {
	data, err = json.Marshal(vault)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (vault *Vault) FindAccountsByURL(url string) []Account {
	var acc []Account
	for _, v := range vault.Accounts {
		if strings.Contains(v.Url, url) {
			acc = append(acc, v)
		}
	}
	return acc
}
func (vault *Vault) DeleteAccountByUrl(url string) bool {
	isDeleted := false
	var acc []Account
	for _, v := range vault.Accounts {
		isMatched := strings.Contains(v.Url, url)
		if !isMatched {
			acc = append(acc, v)
			continue
		}
		isDeleted = true
	}
	vault.Accounts = acc
	vault.save()
	return isDeleted
}

func (vault *Vault) save() {
	vault.UpdatedAt = time.Now()
	data, err := vault.ToBytes()
	if err != nil {
		color.Red(err.Error())
	}
	files.WriteToFile(data, "data.json")
}
