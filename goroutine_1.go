package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i <= 200; i++ {
		go func() {
			time.Sleep(20 * time.Millisecond)
			fmt.Println("buffy is the best")
		}()
	}
}
