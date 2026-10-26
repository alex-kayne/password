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

func NewAccount(urlString, login, password string, passwordLength int) (*Account, error) {
	parsed, err := url.Parse(urlString)
	if err != nil {
		return nil, fmt.Errorf("ошибка в урле %w", err)
	}

	if login == "" {
		return nil, fmt.Errorf("логин не может быть пустым")
	}

	if parsed.Scheme != "http" && parsed.Scheme != "https" {
		return nil, fmt.Errorf("недопустимая схема запроса")
	}

	if parsed.Host == "" {
		return nil, fmt.Errorf("отсутствует хост")
	}

	account := &Account{
		URL:      urlString,
		Login:    login,
		Password: password,
	}

	if password == "" {
		if passwordLength < 1 || passwordLength > 32 {
			return nil, fmt.Errorf("длина пароля должна быть от 1 до 32")
		}

		account.GeneratePassword(passwordLength)
	}

	if password == "" {
		account.GeneratePassword(passwordLength)
	}

	return account, nil
}

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

}
