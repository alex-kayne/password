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
		2: (*account.VaultWithDB).FindAccountByUrl,
		3: (*account.VaultWithDB).FindAccountByLogin,
		4: (*account.VaultWithDB).DeleteAccount,
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
		case 1, 2, 3:
			err = menuMap[num](vault)
		case 5:
			return
		default:
			fmt.Println("Неизвестный параметр выбора")
		}

		if err != nil {
			output.PrintError(err)
		}
	}
}
