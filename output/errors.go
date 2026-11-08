package output

import (
	"github.com/fatih/color"
)

func PrintError(value any) {
	switch v := value.(type) {
	case error:
		color.Red("Ошибка типа error - %s \n", v.Error())
	case string:
		color.Red("Ошибка - %s \n", v)
	case int:
		color.Red("Код ошибки - %d \n", v)
	default:
		color.Red("Все остальные значения %v \n", v)
	}

}
