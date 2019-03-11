package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		go func() {
			wg.Add(1)
			fmt.Println("doing work...")
			time.Sleep(10 * time.Millisecond)
			fmt.Println("work done!")
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("exiting")
}