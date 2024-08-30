package task20

import (
	"fmt"
	"strings"
)

func main() {
	text := "one two three"
	reversedText := reverseWords(text)
	fmt.Println(reversedText)
}

// используем сепаратор в виде пробела для разбивки на слайс из строк
func reverseWords(text string) string {
	words := strings.Split(text, " ")
	//аналогично перевороту строки, меняем элементы слайса местами
	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}

	return strings.Join(words, " ") //собираем строку обратно используя всё тот же сепаратор
}
