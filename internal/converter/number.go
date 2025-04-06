package converter

import (
	"fmt"
	"strconv"
)

func ConvertNumber(input string) string {
	num, err := strconv.Atoi(input)
	if err != nil {
		return "Ошибка: неверное число"
	}

	bin := fmt.Sprintf("%b", num)
	hex := fmt.Sprintf("%X", num)

	return fmt.Sprintf("Двоичная: %s\nШестнадцатеричная: %s", bin, hex)
}
