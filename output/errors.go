package output

import (
	"github.com/fatih/color"
)

func PrintError(value any) {
	vError, ok := value.(error)
	if ok {
		color.Red(vError.Error())
		return
	}

	vString, ok := value.(string)
	if ok {
		color.Red(vString)
		return
	}

	vInt, ok := value.(int)
	if ok {
		color.Red("%d", vInt)
		return
	}

	color.Red("Неподдреживаемый тип")
	return
}
