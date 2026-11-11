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

func CheckUrl(account *Account, url string) bool {
	return strings.Contains(account.URL, url)
}

func CheckLogin(account *Account, login string) bool {
	return strings.Contains(account.Login, login)
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
