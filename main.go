package main

import (
	"demo/password/account"
	"demo/password/menu"
	"fmt"

	"github.com/fatih/color"
)

func main() {
	fmt.Println("Добро пожаловать в менеджер паролей!")
	var urlString, login, password string
	var passwordLength int

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

			err = account.CreateAccount(urlString, login, password, passwordLength)
			if err != nil {
				color.Red("%v", err)
			} else {
				color.Green("Аккаунт выгружен в файл data.json")
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
