package main

import "fmt"
import "strconv"
import "math"

func main() {
	var x string
	for {
		fmt.Println("Please, introduce a float number or 'quit' for exit: ")
		fmt.Scan(&x)
		if x == "quit" {
			break
		} else {
			var y float64
			y, _ = strconv.ParseFloat(x, 64)
			fmt.Printf("Number truncated: %v\n", math.Round(y))
		}
	}
}