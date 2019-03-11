package main

import "fmt"

type Animal interface {
	Bark()
}

type dog struct {
	name string
}

func (dog) Bark() {
	fmt.Println("Voof!")
}

func NewDog(name string) *dog {
	if name == "" {
		return nil
	}
	return &dog{name}
}


func main() {
	var animal Animal

	if animal = NewDog("rufus"); animal != nil {
		animal.Bark()
	}

	if animal = NewDog(""); animal != nil {
		animal.Bark()
	}
}