package task6

import (
	"fmt"
	"sync"
	"time"
)

func byWG() {
	wg := &sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("Goroutine working...")
	}()
	time.Sleep(5 * time.Second)
	wg.Wait()
	fmt.Println("Goroutine done!")
}
