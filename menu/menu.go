package menu

import "fmt"

func GetMenu() {
	text := `Выберите вариант:

1. Создать аккаунт
2. Найти аккаунт
3. Удалить аккаунт
4. Выход

Ваш выбор:`

	fmt.Println(text)
}
