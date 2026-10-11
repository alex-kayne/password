package main

import "fmt"

type Password struct {
	Service  string
	Login    string
	Password string
}

func (p *Password) ChangeLogin(login string) {
	p.Login = login
}

func (p Password) GetDescription() string {
	return p.Service + " " + p.Login
}

func (p *Password) IsEmpty() bool {
	return p.Password == ""
}

func main() {
	yandex := Password{
		Service:  "yandex",
		Login:    "yandex",
		Password: "yandex",
	}

	yandex.ChangeLogin("new_yandex")

	fmt.Println(yandex)

	fmt.Println(yandex.GetDescription())

	fmt.Println(yandex.IsEmpty())

}
