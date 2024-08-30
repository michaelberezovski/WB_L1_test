package task11

import "fmt"

func main() {
	// 2 слайса строк
	one := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	two := []string{"seven", "eleven", "five", "two", "twenty", "sixty", "ninety"}

	setOne := createMap(one)
	setTwo := createMap(two)
	result := intersection(setOne, setTwo)

	for key := range result {
		fmt.Printf("%s ", key)
	}

}

// создаём мапу со значениями
func createMap(array []string) map[string]bool {
	result := make(map[string]bool)
	for _, i := range array {
		result[i] = true
	}
	return result
}

// функция с возвратом новой мапы с общими элементами
func intersection(setOne, setTwo map[string]bool) map[string]bool {
	set := make(map[string]bool)
	for key := range setOne { //в цикле смотрим содержит ли второе множество ключ первого
		if setTwo[key] { //в случае если да - добавляем в результат
			set[key] = true
		}
	}
	return set
}
