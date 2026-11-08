package output

import (
	"github.com/fatih/color"
)

func PrintError(value any) {
	switch v := value.(type) {
	case error:

		color.Red("Ошибка типа error %v \n", v)
	case string:
		color.Red("Ошибка строка %s \n", v)
	default:
		color.Red("Все остальные значения %v \n", v)
	}

}
