package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0 * time.Second; i < 5; i++ {
		go func() {
			time.Sleep(i * time.Second)
			fmt.Println(i)
		}()
	}

	time.Sleep(6 * time.Second)
}
