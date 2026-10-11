package main

import "fmt"

type Password struct {
	Service  string
	Login    string
	Password string
}

func changeLogin(p Password, login string) {
	p.Login = login
}

func changeCurrentLogin(p *Password, login string) {
	p.Login = login
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

	newYandex := yandex

	changeLogin(yandex, "new_yandex")
	changeCurrentLogin(&newYandex, "new_yandex")

	passwords := []Password{yandex, google, newYandex}

	passwords[0].Login = "new login"

	for inx := range passwords {
		passwords[inx].Password = passwords[inx].Password + passwords[inx].Password
	}

	fmt.Printf("%+v\n", passwords)
}
