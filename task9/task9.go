package task9

import "fmt"

func main() {

	array := [5]int{1, 2, 3, 4, 5}

	load := toFirstChannel(array) //пишем данные в первый канал

	result := toSecondChannel(load) // во второй

	for stdout := range result { //выводим
		fmt.Println(stdout)
	}
}

// две функции для записи в каналы
func toFirstChannel(myArray [5]int) <-chan int {
	ch := make(chan int) //создали канал
	go func() {          //запустили горутину
		for _, v := range myArray { //записали массив  в канал
			ch <- v
		}
		defer close(ch) //закрыли канал после
	}()
	return ch
}

func toSecondChannel(load <-chan int) <-chan int {
	ch := make(chan int)
	go func() {
		for v := range load {
			ch <- v * v
		}
		defer close(ch)
	}()
	return ch
}
