package account

import (
	"encoding/json"
	"os"
	"time"

	"github.com/fatih/color"
)

type Vault struct {
	Accounts  []*Account `json:"accounts"`
	UpdatedAt time.Time  `json:"updated_at"`
}

func NewVault() Vault {
	f, err := os.ReadFile("data.json")
	if err != nil {
		color.Red("Не существует файла с названием: %s", "data.json")
	}

	var data Vault

	err = json.Unmarshal(f, &data)

	if err != nil {
		color.Red("Ошибка сериализации: %v", err)
	}

	return data
}

func (v *Vault) AddAccount(account *Account) {
	v.Accounts = append(v.Accounts, account)
	v.UpdatedAt = time.Now()
	data, err := json.Marshal(v)
	if err != nil {
		color.Red("Не удалось сериализовать account")
	}
	err = os.WriteFile("data.json", data, 64)

	if err != nil {
		color.Red("Не удалось сохранить аккаунты %v", err)
	}

}

func (v *Vault) ToBytes() ([]byte, error) {
	bytes, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}
