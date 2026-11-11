package account

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/fatih/color"
)

type ByteReader interface {
	Read() ([]byte, error)
}

type ByteWriter interface {
	Write([]byte) error
}

type Storage interface {
	ByteReader
	ByteWriter
}

type VaultWithDB struct {
	Vault
	db Storage
}

func NewVaultWithDB(db Storage) (*VaultWithDB, error) {
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

func (v *VaultWithDB) CreateAccount() error {
	var urlString, login, password string
	var passwordLength int

	fmt.Println("Введите url, login, password, длину пароля")
	_, err := fmt.Scan(&urlString, &login, &password, &passwordLength)

	if password == "-" {
		password = ""
	}

	if err != nil {
		return err
	}

	newAccount, err := NewAccount(urlString, login, password, passwordLength)
	if err != nil {
		return err
	} else {
		err = v.AddAccount(newAccount)

		if err != nil {
			return err
		}
	}
	return nil
}

func (v *VaultWithDB) FindAccountByUrl() error {
	var urlString string

	fmt.Println("Введите url для поиска account")

	_, err := fmt.Scan(&urlString)

	if err != nil {
		return err
	}

	foundAccounts := v.FindAccounts(urlString, CheckUrl)
	if len(foundAccounts) == 0 {
		fmt.Println("Аккаунтов не найдено.")
	} else {
		for _, foundAc := range foundAccounts {
			foundAc.OutputPassword()
		}
	}

	return nil
}

func (v *VaultWithDB) FindAccountByLogin() error {
	var urlString string

	fmt.Println("Введите login для поиска account")

	_, err := fmt.Scan(&urlString)

	if err != nil {
		return err
	}

	foundAccounts := v.FindAccounts(urlString, CheckLogin)
	if len(foundAccounts) == 0 {
		fmt.Println("Аккаунтов не найдено.")
	} else {
		for _, foundAc := range foundAccounts {
			foundAc.OutputPassword()
		}
	}

	return nil
}

func (v *VaultWithDB) DeleteAccount() error {
	var urlString string

	fmt.Println("Введите url для удаления account")
	_, err := fmt.Scan(&urlString)
	if err != nil {
		return err
	}
	isDeleted, err := v.DeleteAccountByURL(urlString)
	if err != nil {
		return err
	}
	if isDeleted {
		fmt.Println("Аккаунт успешно удален")
	} else {
		fmt.Println("Не удалось удалить аккаунт")
	}
	return nil
}
