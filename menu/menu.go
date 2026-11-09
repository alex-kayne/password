package menu

import (
	"fmt"
)

func GetMenu() (int, error) {
	menuList := []string{
		"Выберите вариант:",
		"1. Создать аккаунт",
		"2. Найти аккаунт",
		"3. Удалить аккаунт",
		"4. Выход",
		"Ваш выбор",
	}

	PrintList(menuList)
	var num int

	_, err := fmt.Scan(&num)

	if err != nil {
		return 0, err
	}
	return num, nil
}

func PrintList[T any](items []T) {
	for inx, item := range items {
		if inx != len(items)-1 {
			fmt.Println(item)
		} else {
			fmt.Printf("%v: ", item)
		}
	}
}
