package menu

import (
	"fmt"
)

func GetMenu() (int, error) {
	text := `Выберите вариант:

1. Создать аккаунт
2. Найти аккаунт
3. Удалить аккаунт
4. Выход

Ваш выбор:`

	fmt.Println(text)
	var num int

	_, err := fmt.Scan(&num)

	if err != nil {
		return 0, err
	}
	return num, nil
}
