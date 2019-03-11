package main

import "fmt"

func main() {
	data := []int{1, 1, 2, 3, 5, 8, 13, 21}
	for i := range data {
		fmt.Printf("data point: %d\n", i)
	}
}
