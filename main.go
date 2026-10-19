package main

import "fmt"
import "math/rand/v2"

var SYMBOLS = []rune{'!', '@', '#', '$', '%', '^', '&', '*', '_', '-', '+', '=', '?', '~', 'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z', 'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z'}

func genPassword(n int) string {
	result := make([]rune, n)
	for inx := range n {
		symInx := rand.IntN(66)
		result[inx] = SYMBOLS[symInx]
	}
	return string(result)
}

func main() {
	fmt.Println("Введите длину пароля от 1 до 32 символов")
	var n int
	_, err := fmt.Scanln(&n)
	if err != nil {
		fmt.Println("Что-то пошло не так")
	}

	if n < 1 || n > 32 {
		fmt.Println("Длина пароля может быть от 0 до 32")
	}

	fmt.Println(genPassword(n))
}
