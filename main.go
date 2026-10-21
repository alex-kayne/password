package main

import "fmt"
import "math/rand/v2"
import "net/url"

type Account struct {
	URL      string
	Login    string
	Password string
}

var allowedSymbols = []rune{'!', '@', '#', '$', '%', '^', '&', '*', '_', '-', '+', '=', '?', '~', 'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z', 'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z', '0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}

func (a *Account) GeneratePassword(n int) {
	result := make([]rune, n)
	for inx := range n {
		symInx := rand.IntN(len(allowedSymbols))
		result[inx] = allowedSymbols[symInx]
	}
	a.Password = string(result)
}

func (a *Account) ChangeLogin(login string) {
	a.Login = login
}

func (a Account) GetDescription() string {
	return a.URL + " " + a.Login
}

func (a *Account) IsEmpty() bool {
	return a.Password == ""
}

func (a *Account) OutputPassword() {
	fmt.Println(a.URL)
	fmt.Println(a.Login)
	fmt.Println(a.Password)
}

func NewAccount(urlString, login, password string) (*Account, error) {
	parsed, err := url.Parse(urlString)

	if err != nil {
		return nil, fmt.Errorf("ошибка в урле %w", err)
	}

	if parsed.Scheme != "http" && parsed.Scheme != "https" {
		return nil, fmt.Errorf("недопустимая схема запроса %w", err)
	}

	if parsed.Host == "" {
		return nil, fmt.Errorf("отсутствует хост %w", err)
	}

	return &Account{
		URL:      urlString,
		Login:    login,
		Password: password,
	}, nil
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

	yandex, err := NewAccount("https://www.google.com/search", "Yandex", "Yandex")

	if err != nil {
		fmt.Println(err)
		return
	}

	yandex.GeneratePassword(n)

	yandex.ChangeLogin("new_yandex")

	fmt.Println(yandex)

	fmt.Println(yandex.GetDescription())

	fmt.Println(yandex.IsEmpty())

	yandex.OutputPassword()

}
