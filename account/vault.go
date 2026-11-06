package account

import (
	"demo/password/files"
	"encoding/json"
	"errors"
	"os"
	"strings"
	"time"

	"github.com/fatih/color"
)

type Vault struct {
	Accounts  []*Account `json:"accounts"`
	UpdatedAt time.Time  `json:"updated_at"`
}

func NewVault() (*Vault, error) {
	f, err := files.ReadFile("data.json")

	if errors.Is(err, os.ErrNotExist) {
		color.Red("Не существует файла с названием: %s", "data.json")
		return &Vault{}, nil
	}

	if err != nil {
		return nil, err
	}

	var vault Vault

	err = json.Unmarshal(f, &vault)

	if err != nil {
		color.Red("Ошибка сериализации: %v", err)
		return nil, err
	}

	return &vault, nil
}

func (v *Vault) Save() error {
	data, err := json.Marshal(v)
	if err != nil {
		color.Red("Не удалось сериализовать account")
		return err
	}
	err = os.WriteFile("data.json", data, 0600)

	if err != nil {
		color.Red("Не удалось сохранить аккаунты %v", err)
		return err
	}

	return nil
}

func (v *Vault) AddAccount(account *Account) error {
	v.Accounts = append(v.Accounts, account)
	v.UpdatedAt = time.Now()
	err := v.Save()
	return err
}

func (v *Vault) ToBytes() ([]byte, error) {
	bytes, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

func (v *Vault) FindAccountsByURL(url string) []*Account {
	var result []*Account
	for _, account := range v.Accounts {
		if strings.Contains(account.URL, url) {
			result = append(result, account)
		}
	}
	return result
}

func (v *Vault) DeleteAccountByURL(url string) (bool, error) {
	var newAccounts []*Account
	isDeleted := false
	for _, account := range v.Accounts {
		if account.URL != url {
			newAccounts = append(newAccounts, account)
		} else {
			isDeleted = true
		}
	}

	if isDeleted {
		v.Accounts = newAccounts
		v.UpdatedAt = time.Now()

		err := v.Save()
		if err != nil {
			return false, err
		}
	}

	return isDeleted, nil
}
