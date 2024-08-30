package task15

import "fmt"

// проблема оригиналього кода в том, что сборщик мусора не будет чистить до момента
//пока есть действующая ссылка
//так что в целях оптимизации стоит убрать ссылку новой выборки на изначальный большой фрагмент текста

var justString string

func createHugeString(length int) string {
	charArray := make([]byte, length)
	for i := range charArray {
		charArray[i] = 'x'
	}
	return string(charArray)
}

func someFunc() {
	v := createHugeString(1 << 10)
	justString = v[:100]
}

func main() {
	someFunc()
	fmt.Println("First 100 characters:", justString)
}
