package main

import (
	"demo/password/account"
	"demo/password/menu"
	"fmt"
	"log"

	"github.com/fatih/color"
)

func main() {
	fmt.Println("Добро пожаловать в менеджер паролей!")
	var urlString, login, password string
	var passwordLength int

	vault, err := account.NewVault()

	if err != nil {
		log.Fatal(err)
	}

	for {
		num, err := menu.GetMenu()

		if err != nil {
			color.Red("%v", err)
			return
		}

		switch num {
		case 1:
			fmt.Println("Введите url, login, password, длину пароля")
			_, err := fmt.Scan(&urlString, &login, &password, &passwordLength)

			if password == "-" {
				password = ""
			}

			if err != nil {
				color.Red("%v", err)
				return
			}

			newAccount, err := account.NewAccount(urlString, login, password, passwordLength)
			if err != nil {
				color.Red("%v", err)
			} else {
				err = vault.AddAccount(newAccount)

				if err != nil {
					color.Red("%v", err)
				}
			}

		case 2:
			fmt.Println("Введите url для поиска account")
			_, err := fmt.Scan(&urlString)
			if err != nil {
				color.Red("%v", err)
				continue
			}
			foundAccounts := vault.FindAccountsByURL(urlString)
			if len(foundAccounts) == 0 {
				fmt.Println("Аккаунтов не найдено.")
			} else {
				for _, foundAc := range foundAccounts {
					foundAc.OutputPassword()
				}
			}

		case 3:
			fmt.Println("Введите url для удаления account")
			_, err := fmt.Scan(&urlString)
			if err != nil {
				color.Red("%v", err)
				continue
			}
			isDeleted, err := vault.DeleteAccountByURL(urlString)
			if err != nil {
				color.Red("%v", err)
				continue
			}
			if isDeleted {
				fmt.Println("Аккаунт успешно удален")
			} else {
				fmt.Println("Не удалось удалить аккаунт")
			}
		case 4:
			return
		default:
			fmt.Println("Неизвестный параметр выбора")
		}
	}
}
