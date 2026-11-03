package main

import (
	"demo/password/account"
	"demo/password/files"
	"demo/password/menu"
	"fmt"

	"github.com/fatih/color"
)

func main() {
	fmt.Println("Добро пожаловать в менеджер паролей!")
	var urlString, login, password string
	var passwordLength int

	vault := account.NewVault()

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
				vault.AddAccount(newAccount)
				data, err := vault.ToBytes()
				if err != nil {
					color.Red("%v", err)
				} else {
					err = files.WriteFile(string(data), "data.json")
					if err != nil {
						color.Red("%v", err)
					} else {
						color.Green("Vault выгружен в файл data.json")
					}
				}
			}

		case 2:
			account.FindAccount()
		case 3:
			account.DeleteAccount()
		case 4:
			return
		default:
			fmt.Println("Неизвестный параметр выбора")
		}
	}
}
