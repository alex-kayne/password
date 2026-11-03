package account

import (
	"encoding/json"
	"time"
)

type Vault struct {
	accounts  []*Account
	updatedAt time.Time `json:"updated_at"`
}

func NewVault() Vault {
	return Vault{}
}

func (v *Vault) AddAccount(account *Account) {
	v.accounts = append(v.accounts, account)
	v.updatedAt = time.Now()
}

func (v *Vault) ToBytes() ([]byte, error) {
	bytes, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}
