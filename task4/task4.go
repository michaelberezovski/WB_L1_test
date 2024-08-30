package task4

import (
	"fmt"
	"time"
)

func main() {
	var Workers int
	fmt.Println("Enter the number of workers ")
	fmt.Scanf("%d\n", &Workers)

	channel := make(chan int) //небуферизованный канал для чтения

	for i := 1; i <= Workers; i++ { //цикл для запуска горутин
		go workerFunc(i, channel)
	}

	for { //цикл для бесконечной записи в кана
		channel <- time.Now().Second()
		time.Sleep(time.Second)
	}
}

func workerFunc(number int, channelToRead <-chan int) { // Функция принимает # воркера и канал
	for {
		timeSec := <-channelToRead
		fmt.Printf("Worker # %d input data: %d\n", number, timeSec)
	}
}
