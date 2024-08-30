package task19

import (
	"fmt"
)

func main() {
	inputString := "abcdefg"
	reversedString := reverseString(inputString)
	fmt.Println(reversedString)
}

// на вход принимаем строку, разбивем ее на слайс из рун, меняем значения местами
func reverseString(s string) string {
	runes := []rune(s)

	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
}
