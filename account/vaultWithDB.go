package account

import (
	"demo/password/db"
	"encoding/json"
	"errors"
	"os"
	"time"

	"github.com/fatih/color"
)

type VaultWithDB struct {
	Vault
	db db.DB
}

func NewVaultWithDB(db db.DB) (*VaultWithDB, error) {
	f, err := db.Read()

	vaultWithDB := VaultWithDB{db: db}

	if errors.Is(err, os.ErrNotExist) {
		color.Red("Не существует файла")
		return &vaultWithDB, nil
	}

	if err != nil {
		return nil, err
	}

	vault, err := NewVault(f)

	if err != nil {
		color.Red("Не удалось создать аккаунты из файла")
		return nil, err
	}

	vaultWithDB.Vault = *vault
	return &vaultWithDB, nil
}

func (v *VaultWithDB) Save() error {
	data, err := json.Marshal(v)
	if err != nil {
		color.Red("Не удалось сериализовать account")
		return err
	}
	err = v.db.Write(data)

	if err != nil {
		color.Red("Не удалось сохранить аккаунты %v", err)
		return err
	}

	return nil
}

func (v *VaultWithDB) AddAccount(account *Account) error {
	v.Accounts = append(v.Accounts, account)
	v.UpdatedAt = time.Now()
	err := v.Save()
	return err
}

func (v *VaultWithDB) DeleteAccountByURL(url string) (bool, error) {
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
