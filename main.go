package main

import (
	"demo/password/account"
	"demo/password/menu"
	"demo/password/output"
	"demo/password/storage/json"
	"fmt"
	"log"
)

func main() {
	fmt.Println("Добро пожаловать в менеджер паролей!")

	jsonDB := json.JSONDB{FilePath: "data.json"}
	vault, err := account.NewVaultWithDB(&jsonDB)

	menuMap := map[int]func(*account.VaultWithDB) error{
		1: (*account.VaultWithDB).CreateAccount,
		2: (*account.VaultWithDB).FindAccount,
		3: (*account.VaultWithDB).DeleteAccount,
	}

	if err != nil {
		log.Fatal(err)
	}

	for {
		num, err := menu.GetMenu()

		if err != nil {
			output.PrintError(err)
			return
		}

		switch num {
		case 1:
			err = menuMap[1](vault)
		case 2:
			err = menuMap[1](vault)
		case 3:
			err = menuMap[1](vault)
		case 4:
			return
		default:
			fmt.Println("Неизвестный параметр выбора")
		}

		if err != nil {
			output.PrintError(err)
		}
	}
}
