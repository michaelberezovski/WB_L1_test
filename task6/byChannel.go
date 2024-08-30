package task6

import (
	"fmt"
	"time"
)

func main() {
	stopChan := make(chan bool) //создаём канал

	go func() { //запускаем горутину
		for {
			select { //обработка возможных кейсов
			case <-stopChan:
				fmt.Println("Goroutine stopping...")
				return
			default:
				fmt.Println("Goroutine working...")
				time.Sleep(1 * time.Second)
			}
		}
	}()

	time.Sleep(5 * time.Second) //создаём ожидание
	close(stopChan)             // останавливаем
	time.Sleep(1 * time.Second) // даём время на установку
	fmt.Println("Done")
}
