package task6

import (
	"context"
	"fmt"
	"time"
)

func byContext() {
	ctx, cancel := context.WithCancel(context.Background())

	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Goroutine stopping...")
				return
			default:
				fmt.Println("Goroutine working...")
				time.Sleep(1 * time.Second)
			}
		}
	}(ctx)

	time.Sleep(5 * time.Second)
	cancel()
	time.Sleep(1 * time.Second)
	fmt.Println("Main function exiting...")
}
