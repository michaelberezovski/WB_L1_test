package task26

import (
	"fmt"
	"strings"
)

func main() {
	testStrings := []string{
		"abcd",
		"abCdefAaf",
		"aabcd",
	}

	for _, str := range testStrings {
		fmt.Printf("Are all characters in \"%s\" unique? %v\n", str, areAllCharactersUnique(str))
	}
}

func areAllCharactersUnique(s string) bool {
	//приводим все значения к нижнему регистру
	s = strings.ToLower(s)
	//добавляем в мапу
	charMap := make(map[rune]bool)
	//при наличии значения в мапе возвращаем false
	for _, char := range s {
		if charMap[char] {
			return false
		}

		charMap[char] = true
	}

	return true
}
