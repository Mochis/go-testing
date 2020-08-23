package main

import (
	"encoding/json"
	"fmt"
)

type Pet struct {
	Name string
	Breed string
	Age int
}

func main() {
	doggy := Pet{"pi", "beagle", 1}
	// Marshal
	doggyJSON, err := json.Marshal(doggy)
	if err != nil {
		panic(err)
	}
	fmt.Println("This is my doggy: ", string(doggyJSON))
	// Unmarshal
	json.Unmarshal(doggyJSON, &doggy)
	fmt.Println("And this is my doggy in Go: ", doggy)
}
