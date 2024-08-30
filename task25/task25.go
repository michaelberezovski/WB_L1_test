package task25

import (
	"fmt"
	"time"
)

//два варианта sleep

// первый хуже, тк грузит CPU, busy waiting loop регулярно проверяет время
func sleepWithForLoop(seconds int) {
	start := time.Now()
	for time.Since(start).Seconds() < float64(seconds) {
	}
}

// второй лучше тем, что time.After создает канал, посылающий сигнал после определенного времени
// а <-sleep блокирует горутину до получения сигнала, соответственно излишне не нагружая проц
func sleepWithChannel(n int) {
	sleep := time.After(time.Duration(n) * time.Second)
	<-sleep
}

func main() {
	fmt.Println("Sleeping for 2 seconds...")
	sleepWithForLoop(2)
	fmt.Println("Awake after 2 seconds!")

	fmt.Println("Sleeping for 3 seconds...")
	sleepWithChannel(3)
	fmt.Println("Awake after 3 seconds!")
}
