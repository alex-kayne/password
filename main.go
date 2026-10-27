package main

import (
	"fmt"
)

func main() {
	var n int

	fmt.Println("Введите длину пароля от 1 до 32 символов")

	_, err := fmt.Scanln(&n)

	if err != nil {
		fmt.Println("ошибка при вводе длины пароля")
		return
	}

	yandex, err := NewAccount("https://www.google.com/search", "Yandex", "", n)

	if err != nil {
		fmt.Println(err)
		return
	}

	yandex.ChangeLogin("new_yandex")

	fmt.Println(yandex)

	fmt.Println(yandex.GetDescription())

	fmt.Println(yandex.IsEmpty())

	yandex.OutputPassword()

	google, err := NewAccountWithTimestamp("https://www.google.com/search", "Google", "", n)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(google.Login)

	google.ChangeLogin("new_google")
}
