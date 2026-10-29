package main

import (
	"demo/password/account"
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/fatih/color"
)

func main() {

	reflectedType := reflect.TypeFor[account.Account]()

	reflectedField, ok := reflectedType.FieldByName("URL")

	if ok {
		fmt.Println(reflectedField.Tag.Get("json"))
	}

	var n int

	fmt.Println("Введите длину пароля от 1 до 32 символов")

	_, err := fmt.Scanln(&n)

	if err != nil {
		color.Red("%v", err)
		return
	}

	yandex, err := account.NewAccount("https://www.google.com/search", "Yandex", "", n)

	if err != nil {
		color.Red("%v", err)
		return
	}

	yandex.ChangeLogin("new_yandex")

	fmt.Println(yandex)

	fmt.Println(yandex.GetDescription())

	fmt.Println(yandex.IsEmpty())

	yandex.OutputPassword()

	google, err := account.NewAccountWithTimestamp("https://www.google.com/search", "Google", "", n)

	if err != nil {
		color.Red("%v", err)
		return
	}

	fmt.Println(google.Login)

	google.ChangeLogin("new_google")

	data, err := json.Marshal(google)
	if err != nil {
		return
	}

	fmt.Println(string(data))
}
