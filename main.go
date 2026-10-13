package main

import "fmt"

type Account struct {
	URL      string
	Login    string
	Password string
}

func (p *Account) ChangeLogin(login string) {
	p.Login = login
}

func (p Account) GetDescription() string {
	return p.URL + " " + p.Login
}

func (p *Account) IsEmpty() bool {
	return p.Password == ""
}

func OutputPassword(a *Account) {
	fmt.Println(a.URL)
	fmt.Println(a.Login)
	fmt.Println(a.Password)
}

func main() {
	yandex := Account{
		URL:      "yandex",
		Login:    "yandex",
		Password: "yandex",
	}

	yandex.ChangeLogin("new_yandex")

	fmt.Println(yandex)

	fmt.Println(yandex.GetDescription())

	fmt.Println(yandex.IsEmpty())

	OutputPassword(&yandex)
}
