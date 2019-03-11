package main

import (
	"encoding/json"
	"fmt"
)

type Data struct {
	username string
	isPrivate bool
}

func main() {
	rawData := `{"username": "buffy", "isPrivate": true}`
	var data Data
	if err := json.Unmarshal([]byte(rawData), &data); err != nil {
		panic(err)
	}
	fmt.Printf("data: %+v\n", data)
}
