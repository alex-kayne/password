package main

import "fmt"

type Password struct {
	Service  string
	Login    string
	Password string
}

func main() {
	yandex := Password{
		Service:  "yandex",
		Login:    "yandex",
		Password: "yandex",
	}

	google := Password{
		Service:  "google",
		Login:    "google",
		Password: "google",
	}

	fmt.Println(yandex)
	fmt.Println(google)

	yandex.Login = "new_yandex"

	newYandex := yandex

	newYandex.Login = "new_new_yandex"

	if yandex.Login == "new_yandex" {
		fmt.Println("Nothing changed")
	}

	passwords := []Password{yandex, google, newYandex}
	for _, password := range passwords {
		fmt.Println(password)
	}
}
