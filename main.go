package main

import "fmt"

type Account struct {
	Url      string
	Login    string
	Password string
}

func (p *Account) ChangeLogin(login string) {
	p.Login = login
}

func (p Account) GetDescription() string {
	return p.Url + " " + p.Login
}

func (p *Account) IsEmpty() bool {
	return p.Password == ""
}

func (a *Account) OutputPassword() {
	fmt.Println(a.Url)
	fmt.Println(a.Login)
	fmt.Println(a.Password)
}

func main() {
	yandex := Account{
		Url:      "yandex",
		Login:    "yandex",
		Password: "yandex",
	}

	yandex.ChangeLogin("new_yandex")

	fmt.Println(yandex)

	fmt.Println(yandex.GetDescription())

	fmt.Println(yandex.IsEmpty())

	yandex.OutputPassword()

}
