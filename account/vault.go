package account

import (
	"encoding/json"
	"strings"
	"time"

	"github.com/fatih/color"
)

type AccountsChecker func(*Account, string) bool
type Vault struct {
	Accounts  []*Account `json:"accounts"`
	UpdatedAt time.Time  `json:"updated_at"`
}

func NewVault(f []byte) (*Vault, error) {
	var vault Vault

	err := json.Unmarshal(f, &vault)

	if err != nil {
		color.Red("Ошибка сериализации: %v", err)
		return nil, err
	}

	return &vault, nil
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

func CheckUrl(account *Account, text string) bool {
	if strings.Contains(account.URL, text) {
		return true
	}
	return false
}

func CheckLogin(account *Account, text string) bool {
	if strings.Contains(account.Login, text) {
		return true
	}
	return false
}

func (v *Vault) FindAccounts(text string, checker AccountsChecker) []*Account {
	var result []*Account
	for _, account := range v.Accounts {
		if checker(account, text) {
			result = append(result, account)
		}
	}
	return result
}
