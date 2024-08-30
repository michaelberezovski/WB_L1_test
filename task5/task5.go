package task5

import (
	"context"
	"fmt"
	"time"
)

const n = 10

func main() {

	workerChan := make(chan int) //создаём канал

	ctx, _ := context.WithTimeout(context.Background(), time.Second*n) //контекст с таймаутом

	go func() { //горутина с записью текущего времени в канал
		for {
			workerChan <- time.Now().Second()
			time.Sleep(time.Second)
		}
	}()

	for { //здесь горутина входит в бесконечный цикл, где в зависимости от типа входящих данных печатает в stdout
		select {
		case <-workerChan:
			fmt.Printf("%d\n", <-workerChan)
		case <-ctx.Done():
			fmt.Println("Time is out")
			return
		}
	}
}
