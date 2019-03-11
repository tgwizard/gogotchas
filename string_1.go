package main

import "fmt"

type MarkedString string

func (m MarkedString) String() string {
	return fmt.Sprintf("<%s>", m)
}

func main() {
	fmt.Printf("%q", MarkedString("foo"))
}
