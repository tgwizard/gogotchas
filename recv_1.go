package main

import "fmt"

type Counter struct {
	count int
}

func (c Counter) Incr() {
	c.count++
}

func (c Counter) Get() int {
	return c.count
}

func main() {
	c := Counter{}
	for i := 0; i < 20; i++ {
		c.Incr()
	}
	fmt.Printf("Count: %d\n", c.Get())
}
