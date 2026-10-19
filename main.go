package main

import "fmt"
import "math/rand/v2"

type Account struct {
	URL      string
	Login    string
	Password string
}

var allowedSymbols = []rune{'!', '@', '#', '$', '%', '^', '&', '*', '_', '-', '+', '=', '?', '~', 'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z', 'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z', '0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}

func genPassword(n int) string {
	result := make([]rune, n)
	for inx := range n {
		symInx := rand.IntN(len(allowedSymbols))
		result[inx] = allowedSymbols[symInx]
	}
	return string(result)
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
	fmt.Println("Введите длину пароля от 1 до 32 символов")
	var n int
	_, err := fmt.Scanln(&n)
	if err != nil {
		fmt.Println("Что-то пошло не так")
		return
	}

	if n < 1 || n > 32 {
		fmt.Println("Длина пароля может быть от 1 до 32")
		return
	}

	yandex := Account{
		URL:      "yandex",
		Login:    "yandex",
		Password: genPassword(n),
	}

	yandex.ChangeLogin("new_yandex")

	fmt.Println(yandex)

	fmt.Println(yandex.GetDescription())

	fmt.Println(yandex.IsEmpty())

	OutputPassword(&yandex)
}
