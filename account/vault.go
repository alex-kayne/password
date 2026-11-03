package account

import (
	"encoding/json"
	"time"
)

type Vault struct {
	Accounts  []*Account `json:"accounts"`
	UpdatedAt time.Time  `json:"updated_at"`
}

func NewVault() Vault {
	return Vault{}
}

func (v *Vault) AddAccount(account *Account) {
	v.Accounts = append(v.Accounts, account)
	v.UpdatedAt = time.Now()
}

func (v *Vault) ToBytes() ([]byte, error) {
	bytes, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}
