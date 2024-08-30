package task12

import "fmt"

func main() {
	a := []string{"cat", "cat", "dog", "cat", "tree"}
	var result []string
	for key := range convert(a) {
		result = append(result, key)
	}
	fmt.Println(result)
}

// функция проходится по слайсу добавляя в результат только уникальные элементы
func convert(array []string) map[string]bool {
	set := make(map[string]bool)
	for _, value := range array {
		set[value] = true
	}
	return set
}
