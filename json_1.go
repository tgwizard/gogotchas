package main

import (
	"encoding/json"
	"fmt"
)

type DataA struct {
	Foo string
}

func main() {
	rawData := `{"Foo": "bar"}`
	var data DataA
	if err := json.Unmarshal([]byte(rawData), data); err != nil {
		panic(err)
	}
	fmt.Printf("data: %+v\n", data)
}
