package main

import (
	"errors"
	"fmt"
	"time"
)

func waitOn(fn func() error, timeout time.Duration) error {
	ch := make(chan error)
	go func() {
		result := fn()
		ch <- result
	}()

	select {
	case result := <-ch:
		return result
	case <-time.After(timeout):
		return errors.New("timeout")
	}
}

func main() {
	result := waitOn(func() error {
		time.Sleep(3 * time.Second)
		return errors.New("no error")
	}, 2*time.Second)

	fmt.Println(result)
}
