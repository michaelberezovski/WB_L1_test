package task7

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var mu sync.Mutex            // для защиты критическо секции
	dataMap := make(map[int]int) // мапа
	var wg sync.WaitGroup        // группа ождания

	numGoroutines := 5 // количество горутин

	for i := 0; i < numGoroutines; i++ { // запуск горутин
		wg.Add(1) // кидаем в счётчик
		go func(id int) {
			defer wg.Done() // убавляем счётчик по завершению
			for j := 0; j < 10; j++ {
				mu.Lock()               // блокируем доступ к критической секции
				dataMap[id*10+j] = id   // пишем в мапу
				mu.Unlock()             // разблокируем
				time.Sleep(time.Second) // создаём имитацию деятельности
			}
		}(i)
	}
	fmt.Println("Goroutines are processing")
	wg.Wait()                          // ждём завершения всех горутин
	fmt.Println("Final map contents:") //печатаем содержимое
	for k, v := range dataMap {
		fmt.Printf("Key: %d, Value: %d\n", k, v)
	}
}
